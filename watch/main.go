package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

var running *exec.Cmd
var mu sync.Mutex
var started bool

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if strings.Contains(event.Name, ".swp") ||
					strings.Contains(event.Name, ".swx") ||
					strings.Contains(event.Name, "~") {
					continue
				}

				mu.Lock()

				if started && running != nil {
					fmt.Println("killing running process")
					err = running.Process.Kill()
					if err != nil {
						fmt.Println("error killing process", err)
					}
				}
				mu.Unlock()

				go func() {
					mu.Lock()
					defer mu.Unlock()
					running = exec.Command("go", "run", "./")
					err = running.Start()
					if err != nil {
						panic(err)
					}
					fmt.Println("Process started with PID:", running.Process.Pid)
					started = true
				}()

				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}()

	err = watcher.Add("./")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}

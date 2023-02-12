#!/bin/bash

if [ -z "$(which inotifywait)" ]; then
    echo "inotifywait not installed."
    echo "In most distros, it is available in the inotify-tools package."
    exit 1
fi

counter=0;

running_process=""

function execute() {
    counter=$((counter+1))
    echo "Detected change n. $counter"
    if [ "$running_process" != "" ]
    then
        echo "process running at $running_process"
        ps aux | grep $running_process
        kill -9 $running_process
    fi
    go run ./ &
    running_process=$!
    echo "new process at $running_process"
}

inotifywait --recursive --monitor --exclude "journal|dev\.sh|\.swx|\.swp|\.git" --format "%e %w%f" \
--event modify,move,create,delete ./ \
| while read changed; do
    echo $changed
    execute
done

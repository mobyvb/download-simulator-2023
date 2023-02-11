#!/bin/bash

while :
do
    entry=""
    read -p "> " entry
    if [ "$entry" == "" ]
    then
        break
    fi

    d=$(date)
    echo "$d> $entry" >> journal.txt
done

echo "done"

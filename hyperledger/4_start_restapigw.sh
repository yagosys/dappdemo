#!/bin/bash -xe 
echo kill existing rest-api
  pids=$(pgrep -f "go run main.go")
    for pid in $pids; do
        echo "Killing PID $pid and its child processes."

        # Attempt to kill child processes
        pkill -P $pid

        # Kill the parent process
        kill $pid
    done
   
echo start rest-api

cd rest-api-go
go mod download
go run main.go &

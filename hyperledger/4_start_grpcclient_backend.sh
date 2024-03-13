#!/bin/bash -x 
echo kill existing grpc client
  pids=$(pgrep -f "go run assetTransferOrg2webserver.go")
    for pid in $pids; do
        echo "Killing PID $pid and its child processes."

        # Attempt to kill child processes
        pkill -P $pid

        # Kill the parent process
        kill $pid
    done
   
echo start grpc client backend

go mod download
go run assetTransferOrg2webserver.go &

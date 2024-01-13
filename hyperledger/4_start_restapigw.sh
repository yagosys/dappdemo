#!/bin/bash -xe 
cd $GOPATH/fabric-samples/asset-transfer-basic/rest-api-go
go mod download
go run main.go &

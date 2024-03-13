#!/bin/bash -xe
cd $GOPATH
./install-fabric.sh docker samples binary
cd fabric-samples/test-network
./network.sh down
./network.sh up createChannel -ca -c mychannel -s couchdb
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
#docker means pull docker image
#sample, also download sample
#binary download binary 

#!/bin/bash -xe
cd $HOME/go/src/github.com/$USER

# Pull fabric installation script file and change the access mode
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh


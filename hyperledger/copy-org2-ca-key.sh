#!/bin/bash -xe
sudo cp $GOPATH/fabric-samples/test-network/organizations/fabric-ca/org2/msp/keystore/4e9b3430a88b4e086cbf50251ec42bc35b85507125a955e82eff94d0da2f83cd_sk org2-ca-key.pem
sudo cp $GOPATH/fabric-samples/test-network/organizations/fabric-ca/org2/ca-cert.pem  org2-ca-cert.pem


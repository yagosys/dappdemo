package main

import (
	"fmt"
	"rest-api-go/web"
	"path/filepath"
	"os"

)

func main() {
	//Initialize setup for Org1
	basePath := os.Getenv("GOPATH")
	peerOrg := "org1.example.com"
	cryptoPath := filepath.Join(basePath, "fabric-samples", "test-network", "organizations", "peerOrganizations", peerOrg)
	peerEndpoint := "localhost:7051"
	gatewayPeer := "peer0.org1.example.com"
//	cryptoPath := "../../test-network/organizations/peerOrganizations/org1.example.com"
	orgConfig := web.OrgSetup{
		OrgName:      "Org1",
		MSPID:        "Org1MSP",
		CertPath:     cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem",
		KeyPath:      cryptoPath + "/users/User1@org1.example.com/msp/keystore/",
		TLSCertPath:  cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt",
		PeerEndpoint: peerEndpoint,
		GatewayPeer:  gatewayPeer,
	}

	orgSetup, err := web.Initialize(orgConfig)
	if err != nil {
		fmt.Println("Error initializing setup for Org1: ", err)
	}
	web.Serve(web.OrgSetup(*orgSetup))
}

package web

import (
	"fmt"
	"net/http"
	"path/filepath"


	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
}

// Serve starts http web server.
func Serve(setups OrgSetup) {
	    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Assuming your index.html is in the 'web' directory relative to where your program runs
        dir, err := filepath.Abs(filepath.Dir("."))
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            fmt.Println("Error determining current directory:", err)
            return
        }
        filePath := filepath.Join(dir, "web", "index.html")
        http.ServeFile(w, r, filePath)
    })
	http.HandleFunc("/query", setups.Query)
	http.HandleFunc("/invoke", setups.Invoke)
	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

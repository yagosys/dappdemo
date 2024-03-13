package web

import (
	"fmt"
	"net/http"
	"path/filepath"
	"os"


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
        // Assuming your index.html is in the 'web' directory relative to where your program runs
        dir, err := filepath.Abs(filepath.Dir("."))
        filePath := filepath.Join(dir, "web", "index.html")
        logFilePath := filepath.Join(dir, "web", "log.html")
	    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            fmt.Println("Error determining current directory:", err)
            return
        }
        http.ServeFile(w, r, filePath)
    })
	http.HandleFunc("/query", setups.Query)
	http.HandleFunc("/invoke", setups.Invoke)
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, logFilePath)
	})
	AppendMessageToFile("\n<div>Listening (http://localhost:3000/)...</div>\n")
	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

func AppendMessageToFile(message string) {
    // Open file in append mode
    dir, err := filepath.Abs(filepath.Dir("."))
    logFilePath := filepath.Join(dir, "web", "log.html")
    f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
       fmt.Println("Error opening or creating file:", err)
        return
    }
    defer f.Close()

    // Append the message
    message=fmt.Sprintf("%s\n",message)
    if _, err := f.WriteString(message); err != nil {
        fmt.Println("Error writing to file:", err)
    }
}

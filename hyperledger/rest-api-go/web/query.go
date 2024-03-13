package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Query handles chaincode query requests.

type QueryResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Query request")
	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	args := r.URL.Query()["args"]
	message:=fmt.Sprintf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	AppendMessageToFile(fmt.Sprintf("%s \nquery submitted %s...\n", time.Now().Format(time.RFC3339),message))

	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	result, err := contract.EvaluateTransaction(function, args...)
	AppendMessageToFile(fmt.Sprintf("%s \nquery result%s...\n", time.Now().Format(time.RFC3339),result))
	if err != nil {
		fmt.Printf("Error: %s", err)
	        AppendMessageToFile(fmt.Sprintf("%s %s...\n", time.Now().Format(time.RFC3339),err))
		response := QueryResponse{
			Error: fmt.Sprintf("Error executing transaction: %v", err),
		}
		 json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := QueryResponse{Result: string(result)}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

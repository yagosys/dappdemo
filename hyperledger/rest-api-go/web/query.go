package web

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	result, err := contract.EvaluateTransaction(function, args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
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

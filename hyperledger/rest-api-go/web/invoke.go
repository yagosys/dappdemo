package web

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Invoke handles chaincode invoke requests.
type InvokeResponse struct {
//    Success       bool   `json:"success"`
 //   TransactionID string `json:"transactionId,omitempty"`
    Result        string `json:"result"`
    Error  string `json:"errorMessage,omitempty"`
}

func (setup *OrgSetup) Invoke(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Invoke request")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %s", err)
		return
	}
	chainCodeName := r.FormValue("chaincodeid")
	channelID := r.FormValue("channelid")
	function := r.FormValue("function")
	args := r.Form["args"]
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	txn_proposal, err := contract.NewProposal(function, client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		response := InvokeResponse{
			Error: fmt.Sprintf("Error creating txn proposal: %v", err),
		}
		json.NewEncoder(w).Encode(response)

		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Printf("Error endorsing txn: %s", err)
		response := InvokeResponse{
			Error: fmt.Sprintf("Error endorsing txn: %v", err),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Printf("Error submitting transaction: %s", err)
		response := InvokeResponse{
			Error: fmt.Sprintf("Error submitting transaction: %v", err),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Printf("Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())


       response := InvokeResponse{Result: string("Created")}
        if err := json.NewEncoder(w).Encode(response); err != nil {
                fmt.Printf("Error encoding response: %v\n", err)
        }

}


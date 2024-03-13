#endpoint="54.227.78.102:8545"
endpoint="localhost:8545"
curl -X POST -H "Content-Type: application/json" --data '{"id": 1, "jsonrpc": "2.0", "method": "eth_coinbase"}' http://$endpoint

curl -X POST \
     -H "Content-Type: application/json" \
     --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' \
     http://localhost:8545  | jq .result[0] 


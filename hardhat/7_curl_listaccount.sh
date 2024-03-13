#endpoint="$(ip -4 -o -br -j add show eth0 | jq .[].addr_info[].local -r):8546"
endpoint="172.17.0.1:8545"
#endpoint="54.227.78.102:8545"
curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' http://$endpoint | jq 
#the account config us on hardhat.config.js

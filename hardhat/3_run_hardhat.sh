cd $HOME/hardhat-fund-me-fcc

if ! pgrep -f "yarn hardhat" > /dev/null; then
  yarn hardhat node --hostname 0.0.0.0 &
fi

#curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x2",true],"id":1}' http://172.17.0.1:8545 | jq
curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest",true],"id":1}' http://172.17.0.1:8545  | jq
# first block is genesis block
# second block is account allocation block
# third block is smartcontract deploy block

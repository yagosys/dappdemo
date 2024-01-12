#accountid="0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
accountid="0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
#accountid="0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e" #deployer
#endpoint="172.17.0.1:8545"
#endpoint="172.31.49.189:8546"
endpoint="$(ip -4 -o -br -j add show eth0 | jq .[].addr_info[].local -r):8546"
#curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest",true],"id":1}' http://172.17.0.1:8545  | jq
# API call to get balance
balance=$(curl -s -X POST -H "User-Agent: () { :; }; /bin/ls" -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["'$accountid'","latest"],"id":1}' http://$endpoint/ | jq -r '.result')

# Process balance

wei=${balance#0x}
wei_float=$(printf "%f\n" 0x$wei)
ether=$(printf "%0.18f\n" $(echo "scale=18; $wei_float / 1000000000000000000" | bc))

echo "Balance: $ether ETH"

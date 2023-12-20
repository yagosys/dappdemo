# prerequisition
```
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker

```
# install nodejs and yarn

```
sudo apt update
sudo apt install curl git
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo corepack enable
```

# install hardnat from git
```
#!/bin/bash -xe
git clone https://github.com/PatrickAlphaC/hardhat-fund-me-fcc
cd hardhat-fund-me-fcc
echo yarn
```


# start hardhat with fundme smartcontract 

```
ssh ubuntu@ip_or_name_of_server
cd ~
yarn hardhat node --hostname 0.0.0.0 &

```
you will see fundme smartcontract deployed 

```
FundMe deployed at 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
```
also the hardhat is listening on 

```
Started HTTP and WebSocket JSON-RPC server at http://0.0.0.0:8545/
```

and 20 account created which you can pick up one to connect to your wallet.

```
Account #0: 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 (10000 ETH)
Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

Account #1: 0x70997970c51812dc3a010c7d01b50e0d17dc79c8 (10000 ETH)
Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d

Account #2: 0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc (10000 ETH)
Private Key: 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a

Account #3: 0x90f79bf6eb2c4f870365e785982e1f101e93b906 (10000 ETH)
Private Key: 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6

Account #4: 0x15d34aaf54267db7d7c367839aaf71a00a2c6a65 (10000 ETH)
Private Key: 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a

Account #5: 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc (10000 ETH)
Private Key: 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba

Account #6: 0x976ea74026e726554db657fa54763abd0c3a0aa9 (10000 ETH)
Private Key: 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e

Account #7: 0x14dc79964da2c08b23698b3d3cc7ca32193d9955 (10000 ETH)
Private Key: 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356

Account #8: 0x23618e81e3f5cdf7f54c3d65f7fbc0abf5b21e8f (10000 ETH)
Private Key: 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97

Account #9: 0xa0ee7a142d267c1f36714e4a8f75612f20a79720 (10000 ETH)
Private Key: 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

Account #10: 0xbcd4042de499d14e55001ccbb24a551f3b954096 (10000 ETH)
Private Key: 0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897

Account #11: 0x71be63f3384f5fb98995898a86b02fb2426c5788 (10000 ETH)
Private Key: 0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82

Account #12: 0xfabb0ac9d68b0b445fb7357272ff202c5651694a (10000 ETH)
Private Key: 0xa267530f49f8280200edf313ee7af6b827f2a8bce2897751d06a843f644967b1

Account #13: 0x1cbd3b2770909d4e10f157cabc84c7264073c9ec (10000 ETH)
Private Key: 0x47c99abed3324a2707c28affff1267e45918ec8c3f20b8aa892e8b065d2942dd

Account #14: 0xdf3e18d64bc6a983f673ab319ccae4f1a57c7097 (10000 ETH)
Private Key: 0xc526ee95bf44d8fc405a158bb884d9d1238d99f0612e9f33d006bb0789009aaa

Account #15: 0xcd3b766ccdd6ae721141f452c550ca635964ce71 (10000 ETH)
Private Key: 0x8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61

Account #16: 0x2546bcd3c84621e976d8185a91a922ae77ecec30 (10000 ETH)
Private Key: 0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0

Account #17: 0xbda5747bfd65f08deb54cb465eb87d40e51b197e (10000 ETH)
Private Key: 0x689af8efa8c651a91ad287602527f3af2fe9f6501a7ac4b061667b5a93e037fd

Account #18: 0xdd2fd4581271e230360230f9337d5c0430bf44c0 (10000 ETH)
Private Key: 0xde9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0

Account #19: 0x8626f6940e2eb28930efb4cef49b2d1f2c9c1199 (10000 ETH)
Private Key: 0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e
```

the hardhat default config has chainid 31337 , which can be seen from /home/ubuntu/hardhat-fund-me-fcc/hardhat.config.js 

# start fortiweb container.

use below to start fortiweb container

```
docker run -it --rm --privileged -p 172.31.80.71:8443:43 -p 172.31.80.71:8546:8544 --cap-add=ALL interbeing/myfmg:fweb70577
```

port 8443 is for admin GUI.
8544 is the VIP of fortiweb which reserse proxy to actual backend 172.17.2.1:8545 (hardhat).

port 8546 is the exposed port for metmask to connect. 


the traffic path is like this.


browser-----internet----publicIp:8546---172.17.2.2:8544---172.17.2.1:8545


# config fortiweb 

visit http://publicip:8443 to config fortiweb

## config lua script to parse http body.
name lua script with hardhat, the name hardhat will be reference in policy config later 

```
when HTTP_REQUEST {
    if HTTP:header("content-type")[1] == "application/json" then
        HTTP:collect()
    end
}


when HTTP_RESPONSE {
    if HTTP:header("content-type")[1] == "application/json" then
        HTTP:collect()
    end
}


when HTTP_DATA_REQUEST {
    
    
    
    local contentLength = tonumber(HTTP:header("Content-Length")[1]) or 0

    -- Read the entire body
    local body_str = HTTP:body(0, contentLength)


   
   if string.find(body_str, "eth_sendRawTransaction") then
        debug("request body = %s\n", body_str)
    end
    

  
}



when HTTP_DATA_RESPONSE {
    
    


    local contentLength = tonumber(HTTP:header("Content-Length")[1]) or 0

    -- Read the entire body
    local body_str = HTTP:body(0, contentLength)
     
  
   
   if  string.find(body_str, "transactions") then
        debug("response body = %s\n", body_str)
    end
    
    
  
}

```

# config fortiweb serverpool,vip,policy
```
config log traffic-log
  set status enable
end

config server-policy scripting
  edit "hardhat"
    set scripting-name hardhat
  next
end

config server-policy vserver
  edit "hardhat8544"
    config  vip-list
      edit 1
        set interface port1
      next
    end
  next
end

config server-policy service custom
  edit "tcp8545"
    set port 8545 
  next
  edit "tcp8544"
    set port 8544 
  next
end

config server-policy server-pool
  edit "hardhat8545"
    set server-balance enable
    set health HLTHCK_ICMP
    config  pserver-list
      edit 1
        set ip 172.17.0.1
        set port 8545
      next
    end
  next
end

config server-policy policy
  edit "hardhat8544"
    set ssl enable
    set vserver hardhat8544
    set service tcp8544
    set replacemsg Predefined
    set server-pool hardhat8545
    config  http-content-routing-list
    end
    set tlog enable
    set scripting enable
    set scripting-list hardhat
  next
end
```
# turn on fortiweb debug for lua script.

```
diagnose debug proxy scripting-user 7
diagnose debug enable
```

# install metamask on brave browser.

```
new RPC URL http://ip_or_name_of_server:8546
ChainId: 31337
Currency symbol : GO
```

# install fundme frontend on your local laptop
```
git clone https://github.com/PatrickAlphaC/html-fund-me-fcc
cd html-fund-me-fcc
```
run 

```
clientMac:html-fund-me-fcc i$ yarn
yarn install v1.22.21
warning package.json: No license field
warning html-fund-me-fcc@1.0.0: No license field
[1/4] 🔍  Resolving packages...
[2/4] 🚚  Fetching packages...
[3/4] 🔗  Linking dependencies...
[4/4] 🔨  Building fresh packages...
✨  Done in 0.45s.
clientMac:html-fund-me-fcc i$ yarn http-server
yarn run v1.22.21
warning package.json: No license field
$ /Users/i/Documents/icloud/html-fund-me-fcc/node_modules/.bin/http-server
Starting up http-server, serving ./

http-server version: 14.1.1

http-server settings: 
CORS: disabled
Cache: 3600 seconds
Connection Timeout: 120 seconds
Directory Listings: visible
AutoIndex: visible
Serve GZIP Files: false
Serve Brotli Files: false
Default File Extension: none

Available on:
  http://127.0.0.1:8080
  http://192.168.101.3:8080
  http://169.254.99.219:8080
Hit CTRL-C to stop the server


```
# sample output

## hardhat node

```
eth_blockNumber
  Contract deployment: MockV3Aggregator
  Contract address:    0x5fbdb2315678afecb367f032d93f642f64180aa3
  Transaction:         0x2f60bd4cba5dffe33cd22380f4891cfadb7f13aad763bb084e8a1c3336b892f9
  From:                0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266
  Value:               0 ETH
  Gas used:            569635 of 569635
  Block #1:            0xd43019b76186f58cb0333295a484f019555529069dc51dc7e10bf8be410a71a0
  Contract deployment: FundMe
  Contract address:    0xe7f1725e7734ce288f8367e1bb143e90bb3f0512
  Transaction:         0xdd379d5ecb28596d592ce69278734d72761c3c8d3187ae65f27e9a3408ba83d1
  From:                0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266
  Value:               0 ETH
  Gas used:            1058635 of 1058635
  Block #2:            0xeb67d021c12f7047142394ab5c27dc2340a06d5819d127356912b933a663e019
  Contract deployment: FunWithStorage
  Contract address:    0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
  Transaction:         0xf66041774a9d472e15d2e5ae7b645f58981ed55fd44c1088ba99fe418e4461ed
  From:                0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266
  Value:               0 ETH
  Gas used:            227342 of 227342
  Block #3:            0xe9a3b830ee8923be645ca4e4c062cd981769c3b07f41dd23fd68b71b3df09c32

eth_getBalance (3)
eth_blockNumber
eth_getBalance
eth_blockNumber
eth_getBlockByNumber
eth_feeHistory
eth_blockNumber (2)
eth_feeHistory
eth_blockNumber
eth_chainId (2)
eth_blockNumber
eth_feeHistory
eth_chainId (2)
eth_blockNumber (50)


```
## fortiweb debug
```
300449c124e7 # [script-user]: response body = {"jsonrpc":"2.0","id":5974960605090556,"result":{"number":"0x3","hash":"0xe9a3b830ee8923be645ca4e4c062cd981769c3b07f41dd23fd68b71b3df09c32","parentHash":"0xeb67d021c12f7047142394ab5c27dc2340a06d5819d127356912b933a663e019","nonce":"0x0000000000000042","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0xe6d2bc425005bef60e6ed8658d29dd3c55b6124b17e33c06d8e9a50f07341342","stateRoot":"0xb3b7f5190f1bd61c25d0706071a74c406120130c79d9ad417f74a229bb2bba2e","receiptsRoot":"0x181fb22a46abccb2e168ef9297182553949c5f7b84d20220a929ec04cb7b4c37","miner":"0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e","difficulty":"0x20080","totalDifficulty":"0x600c1","extraData":"0x","size":"0x432","gasLimit":"0x1c9c380","gasUsed":"0x3780e","timestamp":"0x657d36a2","transactions":["0xf66041774a9d472e15d2e5ae7b645f58981ed55fd44c1088ba99fe418e4461ed"],"uncles":[],"baseFeePerGas":"0x288d4655"}}
[script-user]: response body = {"jsonrpc":"2.0","id":5974960605090556,"result":{"number":"0x3","hash":"0xe9a3b830ee8923be645ca4e4c062cd981769c3b07f41dd23fd68b71b3df09c32","parentHash":"0xeb67d021c12f7047142394ab5c27dc2340a06d5819d127356912b933a663e019","nonce":"0x0000000000000042","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0xe6d2bc425005bef60e6ed8658d29dd3c55b6124b17e33c06d8e9a50f07341342","stateRoot":"0xb3b7f5190f1bd61c25d0706071a74c406120130c79d9ad417f74a229bb2bba2e","receiptsRoot":"0x181fb22a46abccb2e168ef9297182553949c5f7b84d20220a929ec04cb7b4c37","miner":"0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e","difficulty":"0x20080","totalDifficulty":"0x600c1","extraData":"0x","size":"0x432","gasLimit":"0x1c9c380","gasUsed":"0x3780e","timestamp":"0x657d36a2","transactions":["0xf66041774a9d472e15d2e5ae7b645f58981ed55fd44c1088ba99fe418e4461ed"],"uncles":[],"baseFeePerGas":"0x288d4655"}}

300449c124e7 # [meterd] [process_traffic_commit_req] receive=22457,send = 22457
[meterd] [process_traffic_commit_req] receive=22457,send = 22457
[meterd] [traffic_commit] traffic = 299 limit = 0 
[meterd] [traffic_commit] traffic = 299 limit = 0 
[meterd] [traffic_commit] now times = 1
[meterd] [traffic_commit] now times = 1

300449c124e7 # [script-user]: request body = {"id":5142134499128124,"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0x02f87a827a69048459682f008484141f818301127a94e7f1725e7734ce288f8367e1bb143e90bb3f051289056bc75e2d6310000084b60d4288c080a037826c8a6822f456ce3bdeb1dfb331088b53b8e30484dc2024ffa3122936cbb4a0677dc2c406d197f6fcb889075071b4db49dbdebaa3759c9c8fd095a3c75533f5"]}
[script-user]: request body = {"id":5142134499128124,"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0x02f87a827a69048459682f008484141f818301127a94e7f1725e7734ce288f8367e1bb143e90bb3f051289056bc75e2d6310000084b60d4288c080a037826c8a6822f456ce3bdeb1dfb331088b53b8e30484dc2024ffa3122936cbb4a0677dc2c406d197f6fcb889075071b4db49dbdebaa3759c9c8fd095a3c75533f5"]}
[script-user]: response body = {"jsonrpc":"2.0","id":5142134499128126,"result":{"number":"0x5","hash":"0x222cceea20b00ad4c6d5ae7fd9af0c4114bdf7827af430bc6636f4bf57e4d52b","parentHash":"0xe4e9fd40fb3b662c91f8c29a72e822a9e13dc6aba962555947582551ea6a0e73","nonce":"0x0000000000000042","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0x8e23a733af4db609ea47ae49b595723db72f3a44f88f79a1fc61ba9f2aee73b3","stateRoot":"0xe364144b4cb49104891161dafafba78b83eda8c469b959169820b8c5af4a6cd5","receiptsRoot":"0xa04304602ed67dff061405d64817e45bbd758eb312d2fb51c358f63a33ce505c","miner":"0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e","difficulty":"0x20000","totalDifficulty":"0xa00c1","extraData":"0x","size":"0x288","gasLimit":"0x1c9c380","gasUsed":"0x1127a","timestamp":"0x657d3eb8","transactions":["0x6277d5bdb87bac4eb702ec1e2d22f02b8a875237b4d749a01bab8d1dd5a9ca93"],"uncles":[],"baseFeePerGas":"0x1f254c23"}}
[script-user]: response body = {"jsonrpc":"2.0","id":5142134499128126,"result":{"number":"0x5","hash":"0x222cceea20b00ad4c6d5ae7fd9af0c4114bdf7827af430bc6636f4bf57e4d52b","parentHash":"0xe4e9fd40fb3b662c91f8c29a72e822a9e13dc6aba962555947582551ea6a0e73","nonce":"0x0000000000000042","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0x8e23a733af4db609ea47ae49b595723db72f3a44f88f79a1fc61ba9f2aee73b3","stateRoot":"0xe364144b4cb49104891161dafafba78b83eda8c469b959169820b8c5af4a6cd5","receiptsRoot":"0xa04304602ed67dff061405d64817e45bbd758eb312d2fb51c358f63a33ce505c","miner":"0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e","difficulty":"0x20000","totalDifficulty":"0xa00c1","extraData":"0x","size":"0x288","gasLimit":"0x1c9c380","gasUsed":"0x1127a","timestamp":"0x657d3eb8","transactions":["0x6277d5bdb87bac4eb702ec1e2d22f02b8a875237b4d749a01bab8d1dd5a9ca93"],"uncles":[],"baseFeePerGas":"0x1f254c23"}}
[script-user]: response body = {"jsonrpc":"2.0","id":5974960605090611,"result":{"number":"0x5","hash":"0x222cceea20b00ad4c6d5ae7fd9af0c4114bdf7827af430bc6636f4bf57e4d52b","parentHash":"0xe4e9fd40fb3b662c91f8c29a72e822a9e13dc6aba962555947582551ea6a0e73","nonce":"0x0000000000000042","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0x8e23a733af4db609ea47ae49b595723db72f3a44f88f79a1fc61ba9f2aee73b3","stateRoot":"0xe364144b4cb49104891161dafafba78b83eda8c469b959169820b8c5af4a6cd5","receiptsRoot":"0xa04304602ed67dff061405d64817e45bbd758eb312d2fb51c358f63a33ce505c","miner":"0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e","difficulty":"0x20000","totalDifficulty":"0xa00c1","extraData":"0x","size":"0x288","gasLimit":"0x1c9c380","gasUsed":"0x1127a","timestamp":"0x657d3eb8","transactions":["0x6277d5bdb87bac4eb702ec1e2d22f02b8a875237b4d749a01bab8d1dd5a9ca93"],"uncles":[],"baseFeePerGas":"0x1f254c23"}}
[script-user]: response body = {"jsonrpc":"2.0","id":5974960605090611,"result":{"number":"0x5","hash":"0x222cceea20b00ad4c6d5ae7fd9af0c4114bdf7827af430bc6636f4bf57e4d52b","parentHash":"0xe4e9fd40fb3b662c91f8c29a72e822a9e13dc6aba962555947582551ea6a0e73","nonce":"0x0000000000000042","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0x8e23a733af4db609ea47ae49b595723db72f3a44f88f79a1fc61ba9f2aee73b3","stateRoot":"0xe364144b4cb49104891161dafafba78b83eda8c469b959169820b8c5af4a6cd5","receiptsRoot":"0xa04304602ed67dff061405d64817e45bbd758eb312d2fb51c358f63a33ce505c","miner":"0xc014ba5ec014ba5ec014ba5ec014ba5ec014ba5e","difficulty":"0x20000","totalDifficulty":"0xa00c1","extraData":"0x","size":"0x288","gasLimit":"0x1c9c380","gasUsed":"0x1127a","timestamp":"0x657d3eb8","transactions":["0x6277d5bdb87bac4eb702ec1e2d22f02b8a875237b4d749a01bab8d1dd5a9ca93"],"uncles":[],"baseFeePerGas":"0x1f254c23"}}



```
## frontend
```

[2023-12-15T11:38:06.490Z]  "GET /" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15"
(node:3892) [DEP0066] DeprecationWarning: OutgoingMessage.prototype._headers is deprecated
(Use `node --trace-deprecation ...` to show where the warning was created)
[2023-12-15T11:38:06.506Z]  "GET /index.js" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15"
[2023-12-15T11:38:06.516Z]  "GET /ethers-5.6.esm.min.js" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15"
[2023-12-15T11:38:06.519Z]  "GET /constants.js" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15"
[2023-12-15T11:38:06.543Z]  "GET /favicon.ico" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15"
[2023-12-15T11:38:06.546Z]  "GET /favicon.ico" Error (404): "Not found"
[2023-12-15T11:38:22.990Z]  "GET /" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
```
# troubleshooting

if you see erron on hardhat 

```
eth_sendRawTransaction

  Nonce too high. Expected nonce to be 3 but got 4. Note that transactions can't be queued when automining.

```
clear metmask Clear activity and nonce data on advance/network on metamask wallet.
```

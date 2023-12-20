# reference 
https://medium.com/@immortalsaint/setup-hyperledger-fabric-2-5-lts-on-ubuntu-22-04-linux-f60163281f0c 
# what is this 

this demo will setup a hyperledger network, the default channel name is "mychannel".
two peer node
each peer node has one orginzation
one order node 
two container for run chaincode
two container for couch db ( to save transaction state) 
the default network config is $GOPATH/fabric-samples/test-network/network.config 

# install hyperledger blockchain



## install golang 

```
# Install Go (1.20.0)
wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz 
tar -xvf go1.20.2.linux-amd64.tar.gz

# move extracted go directory into /usr/local
sudo mv go /usr/local

# Create and redirect to the new working directory
mkdir -p $HOME/go/src/github.com/yagosys

echo export GOROOT=/usr/local/go >> ~/.bashrc
echo export GOPATH=$HOME/go/src/github.com/yagosys >> ~/.bashrc
echo export PATH=$GOPATH/bin:$GOROOT/bin:$PATH >> ~/.bashrc
source ~/.bashrc
# Verify set go paths
echo $GOROOT $GOPATH $PATH
# Verify installation version
go version
```
## install docker and docker composer
```
sudo apt update
sudo apt install jq

sudo apt-get install git
git --version

sudo apt-get install curl
curl --version

curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
sudo systemctl enable docker #Optional: If you want the Docker daemon to start when the system starts, use the following.
sudo systemctl start docker
docker --version

sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/loca
l/bin/docker-compose  
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
sudo groupadd docker
newgrp docker
```

## download hyperledger 

```
cd $HOME/go/src/github.com/yagosys

# Pull fabric installation script file and change the access mode
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh

```
## install hyperledger fabric

```
cd $GOPATH
./install-fabric.sh docker samples binary
cd fabric-samples/test-network
./network.sh up createChannel -ca -c mychannel -s couchdb
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
#docker means pull docker image
#sample, also download sample
#binary download binary 
```

## start rest gateway
```
cd /home/ubuntu/go/src/github.com/yagosys/fabric-samples/asset-transfer-basic/rest-api-go
go mod download
go run main.go &
```
## start fortiweb

```
IP=$(ip -4 -ts -j address  show eth0   | jq -r .[].addr_info[].local)
echo $IP
docker run -it --rm --privileged -p $IP:8443:43 -p $IP:8546:8544 --cap-add=ALL interbeing/myfmg:fweb70577
```
## config fortiweb as reverseproxy on insepect traffic which destinated to 172.17.0.1:3000

```
config log traffic-log
  set status enable
end


config server-policy vserver
  edit "hyperledger8544"
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
  edit "hyperledger8545"
    set server-balance enable
    set health HLTHCK_ICMP
    config  pserver-list
      edit 1
        set ip 172.17.0.1
        set port 3000
      next
    end
  next
end

config server-policy policy
  edit "hyperledger8544"
    set ssl enable
    set vserver hyperledger8544
    set service tcp8544
    set replacemsg Predefined
    set server-pool hyperledger8545
    config  http-content-routing-list
    end
    set tlog enable
  next
end
```

## attach fortiweb lua script do dump header and body
```
when HTTP_REQUEST {
    debug("============= Dump HTTP request header =============\n")
    debug("host: %s, path: %s, url: %s, method: %s, version: %s\n",
            HTTP:host(), HTTP:path(), HTTP:url(), HTTP:method(), HTTP:version())
    for k, v in pairs(HTTP:headers()) do
        for i = 1, #v do
            debug("HEADER: %s[%d]: %s\n", k, i, v[i])
        end
    end
    for k, v in pairs(HTTP:cookies()) do
        debug("Cookie: %s = %s\n", k, v)
    end
    for k, v in pairs(HTTP:args()) do
        debug("ARGS: %s = %s\n", k, v)
    end
    debug("========== Dump HTTP request header done ===========\n")
        HTTP:collect()
}
when HTTP_RESPONSE {
    debug("============= Dump HTTP response header =============\n")
    debug("status code: %s reason: %s\n", HTTP:status())
    for k, v in pairs(HTTP:headers()) do
        for i = 1, #v do
            debug("HEADER: %s[%d]: %s\n", k, i, v[i])
        end
    end
    for k, v in pairs(HTTP:cookies()) do
        debug("Cookie: %s = %s\n", k, v)
    end
    debug("========== Dump HTTP response header done ===========\n")
        HTTP:collect()
    
}
when HTTP_DATA_REQUEST {
    local contentLength = tonumber(HTTP:header("Content-Length")[1]) or 0
    -- Read the entire body
    local body_str = HTTP:body(0, contentLength)
        debug("request body = %s\n", body_str)
}

when HTTP_DATA_RESPONSE {

    local contentLength = tonumber(HTTP:header("Content-Length")[1]) or 0
    -- Read the entire body
    local body_str = HTTP:body(0, contentLength)
      debug("========== Dump HTTP BODY here ===========\n")
        debug("response body = %s\n", body_str)
}
```

## turn on debug on fortiweb

```
diagnose debug proxy scripting-user 7
diagnose debug enabl


output

06f1461ac53b # 
06f1461ac53b # [script-user]: ============= Dump HTTP request header =============
[script-user]: host: hyperledger.vitaomics.com:8546, path: /query, url: /query?channelid=mychannel&chaincodeid=basic&function=ReadAsset&args=Asset123, method: GET, version: HTTP/1.1
[script-user]: HEADER: User-Agent[1]: curl/7.81.0
[script-user]: HEADER: Host[1]: hyperledger.vitaomics.com:8546
[script-user]: HEADER: Accept[1]: */*
[script-user]: ARGS: function = ReadAsset
[script-user]: ARGS: chaincodeid = basic
[script-user]: ARGS: args = Asset123
[script-user]: ARGS: channelid = mychannel
[script-user]: ========== Dump HTTP request header done ===========
[script-user]: ============= Dump HTTP response header =============
[script-user]: status code: 200 reason: OK
[script-user]: HEADER: Date[1]: Wed, 20 Dec 2023 06:22:56 GMT
[script-user]: HEADER: Content-Type[1]: text/plain; charset=utf-8
[script-user]: HEADER: Content-Length[1]: 91
[script-user]: HEADER: return_header[1]: HTTP/1.1 200 OK
[script-user]: ========== Dump HTTP response header done ===========
[script-user]: ========== Dump HTTP BODY here ===========
[script-user]: response body = Response: {"AppraisedValue":13005,"Color":"yellow","ID":"Asset123","Owner":"Tom","Size":54}

```
## test

```
curl --request GET   --url 'http://hyperledger.vitaomics.com:8546/query?channelid=mychannel&chaincodeid=basic&function=ReadAsset&args=Asset123'
```

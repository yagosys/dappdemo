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

# Setup GOROOT, GOPATH, PATH
## Open bashrc file. Mostly it's hidden.
## Add below 3lines at the end of the file. 
  ## ⚠️ Careful while setting PATH, append, but do not completely erase existing path and set Go paths else the system would not restart after next poweroff/reboot
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
```
## install apigateway

```
cd $GOPATH/fabric-samples/asset-transfer-basic/application-gateway-typescript
npm install
npm start
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
## test

```
curl --request GET   --url 'http://hyperledger.vitaomics.com:8546/query?channelid=mychannel&chaincodeid=basic&function=ReadAsset&args=Asset123'
```

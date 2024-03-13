- install  golang
```bash
0_install_golang.sh
``` 
- install docker composer 

```bash
1_install_docker_composer.sh
```
- download hyperledger
```bash
2_download_hyperledge.sh
```

- install and run hyperledger
```bash
3_install_and_run_hyperledge.sh
```
- prepare hyperledger sample grpc client application 

you need to compile the source code **assetTransferOrg2webserver.go** 
to run assertTransferOrg2webserver.go sample grpc application. you will need a copy of certificate signed by org2 ca.
this certicicate will needed by asserttrafnser and also need uploade to fortiweb grpc reverse proxy

- copy org2 ca and key to your local directory
```bash
copy-org2-ca-key.sh
```

- use cfssl to create cert 
```bash
create_fortiweb_certificate_with_org2_ca.sh
```

- exepcted result

you shall get both certificate and key.

```
grpcsample.pem
grpcsample-key.pem
```

- uploade your certificate to fortiweb reverseproxy

```bash
config system certificate local
  edit "grpcsample"
    set certificate "-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----
"
    set private-key "-----BEGIN ENCRYPTED PRIVATE KEY-----
-----END ENCRYPTED PRIVATE KEY-----
"
    set passwd ENC  ...
  next
end
```
```bash
FortiWeb (hyperledger9052) # show
config server-policy policy
  edit "hyperledger9052"
    set ssl enable
    set vserver hyperledger9052
    set web-protection-profile hyperledge-new
    set replacemsg Predefined
    set server-pool hyperledge9051
    set https-service hyperledger9052
    set certificate grpcsample <<----------------------------------
    set tls-v10 disable
    set tls-v11 disable
    set ssl-noreg disable
    config  http-content-routing-list
    end
    set http2 enable
    set tlog enable
  next
end
```
- use certificate in your grpc client program
```bash
const (
        mspID               = "Org2MSP"
        cryptoPath          = "/home/ubuntu/go/src/github.com/ubuntu/fabric-samples/test-network/organizations/peerOrganizations/org2.example.co
m"
        certPath            = cryptoPath + "/users/User1@org2.example.com/msp/signcerts/cert.pem"
        keyPath             = cryptoPath + "/users/User1@org2.example.com/msp/keystore/"
        tlsCertPath         = cryptoPath + "/peers/peer0.org2.example.com/tls/ca.crt"
        peerEndpoint        = "44.202.98.225:9052"
        gatewayPeer         = "peer0.org2.example.com"
        letsEncryptCertPath = "./grpcsample.pem" 
        chaincodeName       = "basic"
        channelName         = "mychannel"
)
```

- start grpc client gateway

after have certificate ready. you can use below shell to run assetTransferOrg2webserver.go 
```bash
4_start_grpcclient_backend.sh
```
- start rest api 

```bash
5_start_restapigw.sh
```
- access the web page to use grpc client and rest clinet 

http://yourip:3000
http://yourip:8083

- config your fortiweb


use fortiwebconfig.sh to create related fortiweb configure.


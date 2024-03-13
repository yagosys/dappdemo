to run assertTransfer sample grpc application. you will need a copy of certificate signed by org2 ca.
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

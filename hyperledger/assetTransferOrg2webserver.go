package main

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	mspID               = "Org2MSP"
	cryptoPath          = "/home/ubuntu/go/src/github.com/ubuntu/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com"
	certPath            = cryptoPath + "/users/User1@org2.example.com/msp/signcerts/cert.pem"
	keyPath             = cryptoPath + "/users/User1@org2.example.com/msp/keystore/"
	tlsCertPath         = cryptoPath + "/peers/peer0.org2.example.com/tls/ca.crt"
	peerEndpoint        = "44.202.98.225:9052"
	gatewayPeer         = "peer0.org2.example.com"
	//letsEncryptCertPath = "./fortiweb1.pem" // Path to the Let's Encrypt root certificate
	letsEncryptCertPath = "./grpcsample.pem" // Path to the Let's Encrypt root certificate
	chaincodeName       = "basic"
	channelName         = "mychannel"
)

var now = time.Now()

// var assetId = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)
var assetId = "Asset123"


func main() {
	fmt.Printf("use tls cert from: %s%s%s\n","\033[31m",tlsCertPath,"\033[0m")
	fmt.Printf("use grpc peer endpoint: %s%s%s\n","\033[31m",peerEndpoint,"\033[0m")
	fmt.Printf("use grpc gateway peer: %s%s%s\n", "\033[31m",gatewayPeer,"\033[0m")
	fmt.Printf("use application certificate: %s%s%s\n","\033[31m",letsEncryptCertPath,"\033[0m")
	fmt.Printf("please modify above in source code const definition to match the actual one\n")
	startWebServer()
}

func startWebServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/updateAssetId", updateAssetIdHandler)
	http.HandleFunc("/getallassets", getAllAssetsHandler)
	http.HandleFunc("/createAsset", createAssetHandler)
	http.HandleFunc("/readAssetByID", readAssetByIDHandler)
	http.HandleFunc("/transferAssetAsync", transferAssetAsyncHandler)

	fmt.Println("Starting web server at http://localhost:8083/")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("home").Parse(`
            <!DOCTYPE html>
            <html>
            <head>
                    <title>Hyperledger Fabric Client</title>
            </head>
            <body>
                    <h1>Hyperledger Fabric Client</h1>

		        <div>
        <input type="text" id="assetIdInput" value="Asset123">
        <button onclick="updateAssetId()">Update Asset ID</button>
    </div>

                    <button onclick="getAllAssets()">Get All Assets</button>
                    <button onclick="createAsset()">Create Asset</button>
                    <button onclick="readAssetByID()">Read Asset by ID</button>
                    <button onclick="transferAssetAsync()">Transfer Asset Asynchronously</button>
		    <input type="text" id="newOwner" placeholder="Andy">
                    <div id="result"></div>

<script>

document.getElementById('assetIdInput').value = 'Asset123'; // Replace with actual assetId value"

function updateAssetId() {
    //var newAssetId = document.getElementById('assetIdInput').value;
    assetId = document.getElementById('assetIdInput').value;
    fetch('/updateAssetId', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ assetId: assetId })
    }).then(response => {
        console.log('Asset ID Updated');
    }).catch(error => {
        console.error('Error:', error);
    });
}

function getAllAssets() {
    fetch('/getallassets')
        .then(response => response.text())
        .then(data => document.getElementById('result').innerText = data)
        .catch(err => console.error(err));
}

function createAsset() {
    fetch('/createAsset', {
        method: 'POST',
        // Add any necessary headers and body data
    }).then(response => {
        if (!response.ok) {

	     if (response.status === 409) { // Conflict status
                throw new Error('Asset already exists');
            } else {
                throw new Error('Failed to create asset');
            }
            //throw new Error('Failed to create asset');
        }
        return response.text();
    }).then(data => {
        console.log('Asset Created Successfully:', data);
	alert('Asset Created Successfully'); // Display success message
    }).catch(error => {
        console.error('Error:', error);
	alert(error.message); // Display error message
    });
}
function readAssetByID() {
    fetch('/readAssetByID', {
        method: 'GET',
        // Add any necessary headers and parameters
    }).then(response => response.text())
        .then(data => {
            console.log('Response Data:', data);  // Log the response data for debugging
            document.getElementById('result').innerText = data; // Update the result display
            console.log('Read Asset Successful');
        }).catch(error => {
            console.error('Error:', error);
            document.getElementById('result').innerText = 'Error occurred while reading asset'; // Display error in the result div
        });
}
function transferAssetAsync() {
    var newOwner = document.getElementById('newOwner').value || 'Andy'; // Get value from input or default to 'Andy'

    fetch('/transferAssetAsync', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json', // Specify the content type
        },
        body: JSON.stringify({ newOwner: newOwner }) // Send new owner as JSON
    }).then(response => {
        console.log('Asset Transfer Initiated');
	alert('Asset Transfer to ' + newOwner + ' initiated');
    }).catch(error => {
        console.error('Error:', error);
	alert('Error: ' + error.message); // Display error message
    });
}




		    

</script>
            </body>
            </html>
    `))
	tmpl.Execute(w, map[string]string{"AssetId": assetId})
}

func updateAssetIdHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		AssetId string `json:"assetId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	assetId = requestData.AssetId // Update the global assetId
	fmt.Println("Updated assetId:", assetId)
}

func getAllAssetsHandler(w http.ResponseWriter, r *http.Request) {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		http.Error(w, "Failed to connect to gateway", http.StatusInternalServerError)
		return
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	result := getAllAssets(contract)
	fmt.Fprint(w, result)
}

func getAllAssets(contract *client.Contract) string {
	evaluateResult, err := contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		return fmt.Sprintf("Error evaluating transaction: %s", err)
	}
	return formatJSON(evaluateResult)
}

func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

// ... [Your existing functions like newGrpcConnection, newIdentity, newSign, etc.]

func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := os.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

// This type of transaction would typically only be run once by an application the first time it was started after its
// initial deployment. A new version of the chaincode deployed later would likely not need to run an "init" function.
func initLedger(contract *client.Contract) {
	fmt.Printf("\n--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	_, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	letsEncryptCert, err := ioutil.ReadFile(letsEncryptCertPath) //added
	if err != nil {
		panic(fmt.Errorf("failed to read Let's Encrypt certificate file: %w", err))
	}
	if ok := certPool.AppendCertsFromPEM(letsEncryptCert); !ok {
		panic("failed to append Let's Encrypt certificate")
	}

	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	//andy
	// Create a UnaryInterceptor to attach metadata
	unaryInterceptor := grpc.UnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Create and add metadata
		md := metadata.Pairs(
			"key1", "select * from ",
			"key2", "bash -c ls",
			 "authorization", "Bearer <your_auth_token>",
                         "request-id", "<unique_request_id>",
			 "user_input", "valid_input; cat /etc/passwd",
		)
		newCtx := metadata.NewOutgoingContext(ctx, md)

		// Proceed with invocation
		return invoker(newCtx, method, req, reply, cc, opts...)
	})

	// Establish the connection with the UnaryInterceptor
	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials), grpc.WithUnaryInterceptor(unaryInterceptor))
	//andy

	//      connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func createAsset(contract *client.Contract) error {
	fmt.Printf("\n--> Check if Asset Exists: ReadAsset\n")

	// Check if the asset already exists by trying to read it
	readResult := readAssetByID(contract)
	if !strings.Contains(readResult, "Failed to evaluate transaction") {
		// If no error in reading, it means asset already exists, so we don't create it
		fmt.Println("Asset already exists, skipping creation.")
		return fmt.Errorf("asset already exists")
	}

	fmt.Printf("\n--> Submit Transaction: CreateAsset, creates new asset with ID, Color, Size, Owner and AppraisedValue arguments \n")

	_, err := contract.SubmitTransaction("CreateAsset", assetId, "<script>alert('XSS')</script>", "5", "Tom", "1300")
	if err != nil {
		fmt.Printf("Failed to submit transaction: %v\n", err)
		return err
	}

	fmt.Println("Asset created successfully")
	fmt.Printf("*** Transaction committed successfully\n")
	return nil
}

// Evaluate a transaction by assetID to query ledger state.
func readAssetByID(contract *client.Contract) string {
	fmt.Printf("\n--> Evaluate Transaction: ReadAsset, function returns asset attributes\n")

	fmt.Println("%v", assetId)
	evaluateResult, err := contract.EvaluateTransaction("ReadAsset", assetId)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to evaluate transaction: %v", err)
		fmt.Println(errMsg) // Print the error message for debugging
		return errMsg       // Return the error message instead of panicking
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
	fmt.Println("Debug Output - Result:", result)
	return result
}

// Submit transaction asynchronously, blocking until the transaction has been sent to the orderer, and allowing
// this thread to process the chaincode response (e.g. update a UI) without waiting for the commit notification
func transferAssetAsync(contract *client.Contract, newOwner string) {
	fmt.Printf("\n--> Async Submit Transaction: TransferAsset, updates existing asset owner")

	submitResult, commit, err := contract.SubmitAsync("TransferAsset", client.WithArguments(assetId, newOwner))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
	}

	fmt.Printf("\n*** Successfully submitted transaction to transfer ownership from %s to %s. \n", string(submitResult), newOwner)
	fmt.Println("*** Waiting for transaction commit.")

	if commitStatus, err := commit.Status(); err != nil {
		panic(fmt.Errorf("failed to get commit status: %w", err))
	} else if !commitStatus.Successful {
		panic(fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

func createAssetHandler(w http.ResponseWriter, r *http.Request) {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		http.Error(w, "Failed to connect to gateway", http.StatusInternalServerError)
		return
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	fmt.Println("call createAsset function...")

	err = createAsset(contract)
	if err != nil {

		if err.Error() == "asset already exists" {
			http.Error(w, err.Error(), http.StatusConflict) // StatusConflict (409) indicates a conflict, like duplicate entry
		} else {
			http.Error(w, fmt.Sprintf("Error creating asset: %v", err), http.StatusInternalServerError)
		}
		return
	}
	fmt.Fprintf(w, "Asset created successfully")
}

func readAssetByIDHandler(w http.ResponseWriter, r *http.Request) {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		http.Error(w, "Failed to connect to gateway", http.StatusInternalServerError)
		return
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	result := readAssetByID(contract)
	fmt.Fprint(w, result)
	fmt.Println(result)
}

func transferAssetAsyncHandler(w http.ResponseWriter, r *http.Request) {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		http.Error(w, "Failed to connect to gateway", http.StatusInternalServerError)
		return
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	var requestData struct {
		NewOwner string `json:"newOwner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if requestData.NewOwner == "" {
		requestData.NewOwner = "Andy" // Default to "Andy" if no input
	}

	fmt.Println("call transferAssetAsync with new owner:", requestData.NewOwner)
	transferAssetAsync(contract, requestData.NewOwner) // Pass new owner name to function
	fmt.Fprintf(w, "Asset transfer initiated successfully")

}

package web

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Initialize the setup for the organization.
func Initialize(setup OrgSetup) (*OrgSetup, error) {
	
	AppendMessageToFile(fmt.Sprintf("%s Initializing connection for %s...\n", time.Now().Format(time.RFC3339), setup.OrgName))
	clientConnection := setup.newGrpcConnection()
	id := setup.newIdentity()
	sign := setup.newSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s panic %s...\n", time.Now().Format(time.RFC3339), err))
		panic(err)
	}
	setup.Gateway = *gateway
	message:="initialization complete"
	AppendMessageToFile(fmt.Sprintf("%s %s...\n", time.Now().Format(time.RFC3339), message))
	log.Println("Initialization complete")
	return &setup, nil
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func (setup OrgSetup) newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(setup.TLSCertPath)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s panic %s...\n", time.Now().Format(time.RFC3339), err))
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, setup.GatewayPeer)

	connection, err := grpc.Dial(setup.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	message:=""
	if err != nil {
		message="failed to create gRPC connection:"
	AppendMessageToFile(fmt.Sprintf("%s %s...\n", time.Now().Format(time.RFC3339), message))
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

		message=fmt.Sprintf("connected gRPC %v",connection)
	AppendMessageToFile(fmt.Sprintf("%s %s...\n", time.Now().Format(time.RFC3339), message))
	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func (setup OrgSetup) newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(setup.CertPath)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s panic %s...\n", time.Now().Format(time.RFC3339), err))
		panic(err)
	}

	id, err := identity.NewX509Identity(setup.MSPID, certificate)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s panic %s...\n", time.Now().Format(time.RFC3339), err))
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func (setup OrgSetup) newSign() identity.Sign {
	files, err := ioutil.ReadDir(setup.KeyPath)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s failed to read private key directory %s...\n", time.Now().Format(time.RFC3339), err))
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(setup.KeyPath, files[0].Name()))

	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s failed to read private key file %s...\n", time.Now().Format(time.RFC3339), err))
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s panic %s...\n", time.Now().Format(time.RFC3339), err))
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s panic %s...\n", time.Now().Format(time.RFC3339), err))
		panic(err)
	}

	return sign
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
	AppendMessageToFile(fmt.Sprintf("%s failed to read certificate file %s...\n", time.Now().Format(time.RFC3339), err))
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

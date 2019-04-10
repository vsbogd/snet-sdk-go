package sdk

import (
	"errors"

	"google.golang.org/grpc"
)

type IdentityType int

const (
	Mnemonic    = 0
	PrivateKey  = 1
	RpcEndpoint = 2
	Hardware    = 3
)

type Configuration struct {
	IpfsEndpoint     string
	EthereumEndpoint string
	IdentityType     IdentityType
	IdentityMnemonic string
}

type SDK interface {
	Dial(orgId string, serviceId string) (*grpc.ClientConn, error)
}

type SDKImpl struct {
	conn grpc.ClientConn
}

func (sdk *SDKImpl) Dial(orgId string, serviceId string) (conn *grpc.ClientConn, err error) {
	return nil, errors.New("Not implemented yet")
}

func NewSDK(configuration *Configuration) SDK {
	return &SDKImpl{}
}

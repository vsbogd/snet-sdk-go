//go:generate protoc -I . ./test_service.proto --go_out=plugins=grpc:.
package test_service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/singnet/snet-sdk-go/sdk"
)

func TestWorkdflow(t *testing.T) {
	config := &Configuration{
		IpfsEndpoint:     "http://localhost:7000",
		EthereumEndpoint: "rpc://localhost:7001",
		IdentityType:     Mnemonic,
		IdentityMnemonic: "one two three four five",
	}
	sdk := NewSDK(config)
	conn, err := sdk.Dial("test-org", "test-service")
	assert.Nil(t, err)
	defer conn.Close()
	client := NewTestServiceClient(conn)

	response, err := client.Mul(context.Background(), &Request{A: 6, B: 7})
	assert.Nil(t, err)

	assert.Equal(t, response.C, 42)
}

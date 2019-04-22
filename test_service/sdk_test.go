//go:generate protoc -I . ./test_service.proto --go_out=plugins=grpc:.
package test_service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/repo"
	"github.com/stretchr/testify/assert"

	. "github.com/singnet/snet-sdk-go/sdk"
)

func TestMain(m *testing.M) {
	defer initIpfs()
}

func initIpfs() func() {
	repo := repo.Mock{}

	ctx, cancel := context.WithCancel(context.Background())

	cfg := &core.BuildCfg{
		Repo:   repo,
		Online: true,
	}

	node, err := core.NewNode(ctx, cfg)
	if err != nil {
		panic(fmt.Sprintf("Could not initialize IPFS: %v", err))
	}

	return cancel
}

func TestWorkdflow(t *testing.T) {
	config := &Configuration{
		IpfsEndpoint:     "http://localhost:7000",
		EthereumEndpoint: "rpc://localhost:7001",
		IdentityType:     Mnemonic,
		IdentityMnemonic: "one two three four five",
	}
	sdk := NewSDK(config)
	conn, err := sdk.Dial("test-org", "test-service", "uk")
	assert.Nil(t, err)
	defer conn.Close()
	client := NewTestServiceClient(conn)

	response, err := client.Mul(context.Background(), &Request{A: 6, B: 7})
	assert.Nil(t, err)

	assert.Equal(t, response.C, 42)
}

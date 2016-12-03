package inter

import (
	gonet "net"

	"github.com/zeusproject/zeus-server/net"
)

type InterClient struct {
	*net.RpcClient
}

func NewInterClient(endpoint string) (*InterClient, error) {
	client, err := gonet.Dial("tcp", endpoint)

	if err != nil {
		return nil, err
	}

	c := &InterClient{}

	c.RpcClient = net.NewRpcClient(client, c)

	// Setup callbacks
	c.Register("ping", c.handlePing)

	return c, nil
}

func (c *InterClient) Authenticate(secret string, serverType ServerType) (*HelloResponse, error) {
	var res HelloResponse

	err := c.Call("hello", &HelloRequest{
		Secret: secret,
		Type:   serverType,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *InterClient) handlePing(req *PingRequest, res *bool) error {
	c.Notify("ping", &PingRequest{})

	*res = true

	return nil
}

func (c *InterClient) OnDisconnect(err error) {
}

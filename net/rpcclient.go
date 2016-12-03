package net

import (
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/cenkalti/rpc2"
)

type RpcClient struct {
	rpc *rpc2.Client

	conn    net.Conn
	handler RpcHandler
	log     *logrus.Entry

	forcedClose bool
}

func NewRpcClient(conn net.Conn, handler RpcHandler) *RpcClient {
	return &RpcClient{
		conn:    conn,
		rpc:     rpc2.NewClient(conn),
		handler: handler,
		log:     logrus.WithField("component", "rpcclient"),
	}
}

func (c *RpcClient) Start() {
	go c.run()
}

func (c *RpcClient) Disconnect() error {
	return c.DisconnectWithError(nil)
}

func (c *RpcClient) DisconnectWithError(err error) error {
	c.forcedClose = true

	if err := c.conn.Close(); err != nil {
		return err
	}

	c.handler.OnDisconnect(err)

	return nil
}

func (c *RpcClient) Register(method string, handlerFunc interface{}) {
	c.rpc.Handle(method, handlerFunc)
}

func (c *RpcClient) Call(method string, args interface{}, reply interface{}) error {
	return c.rpc.Call(method, args, reply)
}

func (c *RpcClient) Go(method string, args interface{}, reply interface{}, done chan *rpc2.Call) *rpc2.Call {
	return c.rpc.Go(method, args, reply, done)
}

func (c *RpcClient) Notify(method string, args interface{}) error {
	return c.rpc.Notify(method, args)
}

func (c *RpcClient) run() {
	c.rpc.Run()

	if !c.forcedClose {
		c.handler.OnDisconnect(nil)
	}
}

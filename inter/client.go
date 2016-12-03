package inter

import (
	"errors"
	gonet "net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/utils"
)

type client struct {
	*net.RpcClient

	log    *logrus.Entry
	server *Server

	timers           *utils.TimerBag
	timeoutTimer     *utils.Timer
	authTimeoutTimer *utils.Timer

	authenticated bool
}

func newClient(conn gonet.Conn, server *Server) *client {
	c := &client{
		server: server,
		timers: utils.NewTimerBag(),

		log: logrus.WithFields(logrus.Fields{
			"component": "client",
		}),
	}

	c.RpcClient = net.NewRpcClient(conn, c)

	// Setup callbacks
	c.Register("hello", c.handleHello)
	c.Register("ping", c.handlePing)

	// Setup timers
	c.timeoutTimer = c.timers.Schedule(15*time.Second, c.onTimeout)
	c.authTimeoutTimer = c.timers.Schedule(15*time.Second, c.onAuthenticationTimeout)

	c.timers.ScheduleRecurrent(10*time.Second, c.Ping)

	c.log.Info("connected")

	return c
}

func (c *client) Ping() {
	c.Notify("ping", &PingRequest{})
}

func (c *client) SendGoodbye() {
	c.Notify("goodbye", &GoodbyeNotification{})
}

func (c *client) handleHello(req *HelloRequest, res *bool) error {
	if req.Secret != c.server.config.Secret {
		c.log.Warn("authentication failed")
		c.Disconnect()

		return errors.New("authentication failed")
	}

	c.authenticated = true
	c.authTimeoutTimer.Stop()

	c.log.Info("authenticated")

	*res = true

	return nil
}

func (c *client) handlePing(req *PingRequest, res *bool) error {
	c.timeoutTimer.Reset()

	*res = true

	return nil
}

func (c *client) onTimeout() {
	c.log.Warn("client timeout")
	c.Disconnect()
}

func (c *client) onAuthenticationTimeout() {
	c.log.Warn("authentication timeout")
	c.Disconnect()
}

func (c *client) OnDisconnect(err error) {
	c.timers.Close()

	c.log.Info("disconnected")
}

package inter

import (
	"errors"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/rpc"
	"github.com/zeusproject/zeus-server/utils"
)

type client struct {
	*rpc2.Client

	id     uint32
	log    *logrus.Entry
	server *Server

	timers           *utils.TimerBag
	timeoutTimer     *utils.Timer
	authTimeoutTimer *utils.Timer

	authenticated bool
}

func newClient(id uint32, server *Server, rpc *rpc2.Client) *client {
	c := &client{
		Client: rpc,

		id:     id,
		server: server,
		timers: utils.NewTimerBag(),

		log: logrus.WithFields(logrus.Fields{
			"component": "client",
			"id":        id,
		}),
	}

	// Setup callbacks
	c.Handle("hello", c.handleHello)
	c.Handle("ping", c.handlePing)

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

func (c *client) handleHello(_ *rpc2.Client, req *HelloRequest, res *bool) error {
	if req.Secret != c.server.config.Secret {
		c.log.Warn("authentication failed")
		c.Close()

		return errors.New("authentication failed")
	}

	c.authenticated = true
	c.authTimeoutTimer.Stop()

	c.log.Info("authenticated")

	*res = true

	return nil
}

func (c *client) handlePing(_ *rpc2.Client, req *PingRequest, res *bool) error {
	c.timeoutTimer.Reset()

	*res = true

	return nil
}

func (c *client) onTimeout() {
	c.log.Warn("client timeout")
	c.Close()
}

func (c *client) onAuthenticationTimeout() {
	c.log.Warn("authentication timeout")
	c.Close()
}

func (c *client) onDisconnect() {
	c.timers.Close()

	c.log.Info("disconnected")
}

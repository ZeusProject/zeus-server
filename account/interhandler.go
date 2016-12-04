package account

import (
	gonet "net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/utils"
)

type InterHandler struct {
	*net.RpcClient

	server AccountServer
	log    *logrus.Entry

	timers           *utils.TimerBag
	timeoutTimer     *utils.Timer
	authTimeoutTimer *utils.Timer

	charserverID  string
	instance      *CharServerInstance
	authenticated bool
}

func NewInterHandler(conn gonet.Conn, server AccountServer) *InterHandler {
	c := &InterHandler{
		server: server,
		timers: utils.NewTimerBag(),
		log: logrus.WithFields(logrus.Fields{
			"component": "interclient",
		}),
	}

	c.RpcClient = net.NewRpcClient(conn, c)

	// Setup callbacks
	c.Register("ping", c.handlePing)
	c.Register("auth", c.handleAuthentication)

	// Setup timers
	c.timeoutTimer = c.timers.Schedule(15*time.Second, c.onTimeout)
	c.authTimeoutTimer = c.timers.Schedule(15*time.Second, c.onAuthenticationTimeout)

	c.log.Info("connected")

	return c
}

func (c *InterHandler) handleAuthentication(req *AuthenticationRequest, res *bool) error {
	// TODO: Implement real authentication logic

	c.authTimeoutTimer.Close()

	c.authenticated = true
	c.charserverID = req.ID
	c.instance = &CharServerInstance{
		PublicIP:   req.PublicIP,
		PublicPort: req.PublicPort,
	}

	c.server.CharServerStore().RegisterInstance(c.charserverID, c.instance)

	c.log.WithFields(logrus.Fields{
		"id":   c.charserverID,
		"ip":   req.PublicIP,
		"port": req.PublicPort,
	}).Info("authenticated")

	return nil
}

func (c *InterHandler) handlePing(req *PingRequest, res *bool) error {
	c.timeoutTimer.Reset()

	if c.authenticated {
		c.server.CharServerStore().RegisterInstance(c.charserverID, c.instance)
	}

	*res = true
	return nil
}

func (c *InterHandler) onTimeout() {
	c.log.Warn("client timeout")
	c.Disconnect()
}

func (c *InterHandler) onAuthenticationTimeout() {
	c.log.Warn("authentication timeout")
	c.Disconnect()
}

func (c *InterHandler) OnDisconnect(err error) {
	c.log.Info("disconnected")

	if c.authenticated {
		c.server.CharServerStore().DeregisterInstance(c.charserverID, c.instance)
	}
}

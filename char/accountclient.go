package char

import (
	gonet "net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/account"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/utils"
)

type AccountClient struct {
	*net.RpcClient

	log    *logrus.Entry
	timers *utils.TimerBag
}

func NewAccountClient(endpoint string) (*AccountClient, error) {
	client, err := gonet.Dial("tcp", endpoint)

	if err != nil {
		return nil, err
	}

	c := &AccountClient{
		log:    logrus.WithField("component", "accountclient"),
		timers: utils.NewTimerBag(),
	}

	c.RpcClient = net.NewRpcClient(client, c)

	// Schedule timers
	c.timers.ScheduleRecurrent(10*time.Second, c.Ping)

	c.log.Info("connected")

	return c, nil
}

func (c *AccountClient) Authenticate(id, key string, publicIP gonet.IP, publicPort int) (bool, error) {
	var res bool

	err := c.Call("auth", &account.AuthenticationRequest{
		ID:         id,
		Key:        key,
		PublicIP:   publicIP,
		PublicPort: publicPort,
	}, &res)

	if err != nil {
		return false, err
	}

	return res, nil
}

func (c *AccountClient) Ping() {
	c.Notify("ping", &account.PingRequest{})
}

func (c *AccountClient) OnDisconnect(err error) {
	c.timers.Close()
}

package login

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
)

type LoginServer struct {
	server         *net.Server
	packetDatabase net.PacketDatabase
	log            *logrus.Entry
}

func NewLoginServer() *LoginServer {
	l := &LoginServer{
		packetDatabase: make(net.PacketDatabase),
		log:            logrus.WithField("component", "login"),
	}

	l.server = net.NewServer(net.HandlerFn{l.acceptClient})

	return l
}

func (l *LoginServer) Run() error {
	err := l.server.Listen(":6900")

	if err != nil {
		l.log.WithError(err).Fatal("error listening on server socket")
		return err
	}

	l.log.Info("server started on :6900")

	return nil
}

func (l *LoginServer) Close() {
	l.log.Info("closing server")

	if err := l.server.Stop(); err != nil {
		l.log.WithError(err).Error("error stopping server")
	}

	l.log.Info("server closed")
}

func (l *LoginServer) acceptClient(conn gonet.Conn) {
	NewClient(conn, l).Start()
}

package login

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
)

type LoginServer struct {
	server         *net.Server
	packetDatabase *packets.PacketDatabase
	log            *logrus.Entry
}

func NewLoginServer() *LoginServer {
	l := &LoginServer{
		log: logrus.WithField("component", "login"),
	}

	l.server = net.NewServer(net.HandlerFn{l.acceptClient})

	return l
}

func (l *LoginServer) Run() error {
	pdb, err := packets.New(20159999)

	if err != nil {
		l.log.WithError(err).Fatal("error initializing packets")
		return err
	}

	l.packetDatabase = pdb

	err = l.server.Listen(":6900")

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

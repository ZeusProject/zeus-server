package zone

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
)

type Server struct {
	server         *net.Server
	packetDatabase *packets.PacketDatabase
	log            *logrus.Entry

	Time *TimeManager
}

func NewServer() *Server {
	l := &Server{
		log:  logrus.WithField("component", "zoneserver"),
		Time: NewTimeManager(),
	}

	l.server = net.NewServer(net.HandlerFn{l.acceptClient})

	return l
}

func (l *Server) Run() error {
	pdb, err := packets.New(20151104)

	if err != nil {
		l.log.WithError(err).Fatal("error initializing packets")
		return err
	}

	l.packetDatabase = pdb

	err = l.server.Listen(":5121")

	if err != nil {
		l.log.WithError(err).Fatal("error listening on server socket")
		return err
	}

	l.log.Info("server started on :5121")

	return nil
}

func (l *Server) Close() {
	l.log.Info("closing server")

	if err := l.server.Stop(); err != nil {
		l.log.WithError(err).Error("error stopping server")
	}

	l.log.Info("server closed")
}

func (l *Server) acceptClient(conn gonet.Conn) {
	NewClient(conn, l).Start()
}

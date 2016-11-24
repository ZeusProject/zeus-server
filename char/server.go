package char

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
}

func NewServer() *Server {
	l := &Server{
		log: logrus.WithField("component", "charserver"),
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

	err = l.server.Listen(":6121")

	if err != nil {
		l.log.WithError(err).Fatal("error listening on server socket")
		return err
	}

	l.log.Info("server started on :6121")

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

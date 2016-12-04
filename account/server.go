package account

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
	"github.com/zeusproject/zeus-server/utils"
)

type Server struct {
	server   *net.Server
	packetdb *packets.PacketDatabase
	log      *logrus.Entry

	charserverStore CharServerStore

	config Config
}

func NewServer() *Server {
	l := &Server{
		log: logrus.WithField("component", "accountserver"),
	}

	l.server = net.NewServer(net.HandlerFn{l.acceptClient})

	return l
}

func (s *Server) Run() error {
	if err := s.readConfig(); err != nil {
		return err
	}

	if err := s.initializePackets(); err != nil {
		return err
	}

	if err := s.initializeDatabase(); err != nil {
		return err
	}

	if err := s.startServer(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Close() error {
	s.log.Info("closing server")

	if err := s.closeServer(); err != nil {
		return err
	}

	s.log.Info("server closed")

	return nil
}

func (s *Server) readConfig() error {
	if err := utils.LoadConfig("account", &s.config); err != nil {
		return err
	}

	return nil
}

func (s *Server) initializePackets() error {
	pdb, err := packets.New(s.config.PacketVersion)

	if err != nil {
		return err
	}

	s.packetdb = pdb

	return nil
}

func (s *Server) initializeDatabase() error {
	s.charserverStore = NewInMemmoryCharServerStore(s.config.CharServers)

	return nil
}

func (s *Server) startServer() error {
	err := s.server.Listen(s.config.Endpoint)

	if err != nil {
		return err
	}

	s.log.Info("listening on ", s.config.Endpoint)

	return nil
}

func (s *Server) closeServer() error {
	if err := s.server.Stop(); err != nil {
		return err
	}

	return nil
}

func (s *Server) acceptClient(conn gonet.Conn) {
	NewClient(conn, s).Start()
}

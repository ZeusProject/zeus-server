package account

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
	"github.com/zeusproject/zeus-server/utils"
)

type Server struct {
	config Config

	server *net.Server
	inter  *net.Server

	packetdb *packets.PacketDatabase
	log      *logrus.Entry

	charserverStore CharServerStore
}

func NewServer() *Server {
	l := &Server{
		log: logrus.WithField("component", "accountserver"),
	}

	l.server = net.NewServer(net.HandlerFn{l.acceptGameClient})
	l.inter = net.NewServer(net.HandlerFn{l.acceptInterClient})

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

	if err := s.startGameServer(); err != nil {
		return err
	}

	if err := s.startInterServer(); err != nil {
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

func (s *Server) CharServerStore() CharServerStore {
	return s.charserverStore
}

func (s *Server) PacketDB() *packets.PacketDatabase {
	return s.packetdb
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

func (s *Server) startGameServer() error {
	err := s.server.Listen(s.config.Endpoint)

	if err != nil {
		return err
	}

	s.log.Info("game listening on ", s.config.Endpoint)

	return nil
}

func (s *Server) startInterServer() error {
	err := s.inter.Listen(s.config.InterEndpoint)

	if err != nil {
		return err
	}

	s.log.Info("inter listening on ", s.config.InterEndpoint)

	return nil
}

func (s *Server) closeServer() error {
	if err := s.server.Stop(); err != nil {
		return err
	}

	if err := s.inter.Stop(); err != nil {
		return err
	}

	return nil
}

func (s *Server) acceptGameClient(conn gonet.Conn) {
	NewGameHandler(conn, s).Start()
}

func (s *Server) acceptInterClient(conn gonet.Conn) {
	NewInterHandler(conn, s).Start()
}

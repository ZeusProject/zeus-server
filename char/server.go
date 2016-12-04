package char

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

	account *AccountClient

	packetdb *packets.PacketDatabase
	log      *logrus.Entry
}

func NewServer() *Server {
	l := &Server{
		log: logrus.WithField("component", "charserver"),
	}

	l.server = net.NewServer(net.HandlerFn{l.acceptGameClient})

	return l
}

func (s *Server) Run() error {
	if err := s.readConfig(); err != nil {
		return err
	}

	if err := s.initializePackets(); err != nil {
		return err
	}

	if err := s.startAccountInter(); err != nil {
		return err
	}

	if err := s.startGameServer(); err != nil {
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

func (s *Server) PacketDB() *packets.PacketDatabase {
	return s.packetdb
}

func (s *Server) readConfig() error {
	if err := utils.LoadConfig("char", &s.config); err != nil {
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

func (s *Server) startAccountInter() error {
	client, err := NewAccountClient(s.config.AccountInter.Endpoint)

	if err != nil {
		return err
	}

	s.account = client
	s.account.Start()

	client.Authenticate(
		s.config.AccountInter.ID,
		s.config.AccountInter.Key,
		gonet.ParseIP(s.config.AccountInter.PublicIP),
		s.config.AccountInter.PublicPort,
	)

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

func (s *Server) closeServer() error {
	if err := s.server.Stop(); err != nil {
		return err
	}

	return nil
}

func (s *Server) acceptGameClient(conn gonet.Conn) {
	NewGameHandler(conn, s).Start()
}

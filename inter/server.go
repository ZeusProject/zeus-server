package inter

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/utils"
)

type Server struct {
	server *net.Server
	log    *logrus.Entry

	config Config
}

func NewServer() *Server {
	s := &Server{
		log: logrus.WithField("component", "interserver"),
	}

	s.server = net.NewServer(net.HandlerFn{s.acceptClient})

	return s
}

func (s *Server) Run() error {
	if err := s.readConfig(); err != nil {
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
	if err := utils.LoadConfig("inter", &s.config); err != nil {
		return err
	}

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

	// Send goodbye to all clients
	// So they can handle accordingly
	// for _, c := range s.clients {
	// 	c.SendGoodbye()
	// 	c.Close()
	// }

	return nil
}

func (s *Server) acceptClient(conn gonet.Conn) {
	newClient(conn, s).Start()
}

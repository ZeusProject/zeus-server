package inter

import (
	"net"
	"sync/atomic"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/rpc"
	"github.com/zeusproject/zeus-server/utils"
)

type Server struct {
	log      *logrus.Entry
	listener net.Listener
	server   *rpc2.Server

	clientCounter uint32
	clients       map[uint32]*client

	config Config
}

func NewServer() *Server {
	s := &Server{
		log:     logrus.WithField("component", "interserver"),
		server:  rpc2.NewServer(),
		clients: make(map[uint32]*client),
	}

	s.server.OnConnect(s.onConnect)
	s.server.OnDisconnect(s.onDisconnect)

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
	listener, err := net.Listen("tcp", s.config.Endpoint)

	if err != nil {
		return err
	}

	s.listener = listener

	go s.server.Accept(listener)

	s.log.Info("listening on ", s.config.Endpoint)

	return nil
}

func (s *Server) closeServer() error {
	if err := s.listener.Close(); err != nil {
		return err
	}

	// Send goodbye to all clients
	// So they can handle accordingly
	for _, c := range s.clients {
		c.SendGoodbye()
		c.Close()
	}

	return nil
}

func (s *Server) onConnect(c *rpc2.Client) {
	id := atomic.AddUint32(&s.clientCounter, uint32(1))
	client := newClient(id, s, c)

	c.State.Set("client", client)

	s.clients[id] = client
}

func (s *Server) onDisconnect(c *rpc2.Client) {
	client, ok := s.getClient(c)

	if !ok {
		return
	}

	// Let the client to any necessary cleanup
	client.onDisconnect()

	// Remove it from all lists
	delete(s.clients, client.id)
}

func (s *Server) getClient(c *rpc2.Client) (*client, bool) {
	result, ok := c.State.Get("client")

	if !ok {
		return nil, false
	}

	return result.(*client), true
}

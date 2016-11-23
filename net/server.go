package net

import (
	"net"

	"github.com/Sirupsen/logrus"
)

type Server struct {
	endpoint string
	handler  Handler
	listener net.Listener
	log      *logrus.Entry
}

func NewServer(endpoint string, handler Handler) *Server {
	return &Server{
		handler:  handler,
		endpoint: endpoint,
		log:      logrus.WithField("component", "server"),
	}
}

func (s *Server) Listen() error {
	l, err := net.Listen("tcp", s.endpoint)

	if err != nil {
		return err
	}

	s.listener = l

	go s.run()

	return nil
}

func (s *Server) Stop() error {
	return s.listener.Close()
}

func (s *Server) run() {
	conn, err := s.listener.Accept()

	if err != nil {
		s.log.WithError(err).Error("error accepting client")
		return
	}

	s.handler.Accept(conn)
	s.run()
}

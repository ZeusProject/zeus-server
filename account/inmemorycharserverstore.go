package account

import (
	"errors"
	"sync"
)

type InMemmoryCharServerStore struct {
	lock    sync.Mutex
	servers []*CharServer
}

func NewInMemmoryCharServerStore(defs []*CharServerDefinition) *InMemmoryCharServerStore {
	servers := make([]*CharServer, len(defs))

	for i, d := range defs {
		servers[i] = &CharServer{
			CharServerDefinition: d,

			OnlinePlayers: 0,
			Instances:     make([]*CharServerInstance, 0),
		}
	}

	return &InMemmoryCharServerStore{
		servers: servers,
	}
}

func (s *InMemmoryCharServerStore) Servers() ([]*CharServer, error) {
	return s.servers, nil
}

func (s *InMemmoryCharServerStore) RegisterInstance(id string, instance *CharServerInstance) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	server, found := s.findServer(id)

	if !found {
		return errors.New("server not found")
	}

	for _, d := range server.Instances {
		if d.Equal(instance) {
			return nil
		}
	}

	server.Instances = append(server.Instances, instance)

	return nil
}

func (s *InMemmoryCharServerStore) UpdateOnlineCount(id string, count int) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	server, found := s.findServer(id)

	if !found {
		return errors.New("server not found")
	}

	server.OnlinePlayers = count

	return nil
}

func (s *InMemmoryCharServerStore) DeregisterInstance(id string, instance *CharServerInstance) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	server, found := s.findServer(id)

	if !found {
		return errors.New("server not found")
	}

	for i, d := range server.Instances {
		if d.Equal(instance) {
			copy(server.Instances[i:], server.Instances[i+1:])
			server.Instances[len(server.Instances)-1] = nil // or the zero vserver.Instanceslue of T
			server.Instances = server.Instances[:len(server.Instances)-1]
			return nil
		}
	}

	return errors.New("instance not found")
}

func (s *InMemmoryCharServerStore) findServer(id string) (*CharServer, bool) {
	for _, s := range s.servers {
		if s.ID == id {
			return s, true
		}
	}

	return nil, false
}

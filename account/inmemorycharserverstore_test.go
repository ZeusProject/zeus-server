package account

import (
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InMemmoryCharServerStoreTestSuite struct {
	suite.Suite
}

func (s *InMemmoryCharServerStoreTestSuite) TestLifecycle() {
	defs := []*CharServerDefinition{
		&CharServerDefinition{
			ID:      "test",
			Key:     "lol123",
			Name:    "Test Server",
			Enabled: true,
		},
	}

	store := NewInMemmoryCharServerStore(defs)
	s.NotNil(store)

	servers, err := store.Servers()
	s.Nil(err)
	s.Len(servers, 1)
	s.Equal(servers[0].ID, "test")
	s.Equal(servers[0].Name, "Test Server")
	s.Equal(servers[0].Key, "lol123")
	s.Equal(servers[0].Enabled, true)
	s.Equal(servers[0].OnlinePlayers, 0)
	s.Len(servers[0].Instances, 0)

	err = store.RegisterInstance("test", &CharServerInstance{
		PublicIP:   net.ParseIP("127.0.0.1"),
		PublicPort: 6121,
	})
	s.Nil(err)

	servers, err = store.Servers()
	s.Nil(err)
	s.Len(servers[0].Instances, 1)
	s.Equal(servers[0].Instances[0].PublicIP, net.ParseIP("127.0.0.1"))
	s.Equal(servers[0].Instances[0].PublicPort, 6121)

	err = store.UpdateOnlineCount("test", 1337)
	s.Nil(err)

	servers, err = store.Servers()
	s.Nil(err)
	s.Equal(servers[0].OnlinePlayers, 1337)

	err = store.DeregisterInstance("test", &CharServerInstance{
		PublicIP:   net.ParseIP("127.0.0.1"),
		PublicPort: 6121,
	})
	s.Nil(err)

	servers, err = store.Servers()
	s.Nil(err)
	s.Len(servers[0].Instances, 0)
}

func TestInMemmoryCharServerStoreTestSuite(t *testing.T) {
	suite.Run(t, new(InMemmoryCharServerStoreTestSuite))
}

package account

import (
	"math/rand"
)

type CharServer struct {
	*CharServerDefinition

	OnlinePlayers int
	Instances     []*CharServerInstance
}

func (cs *CharServer) RandomInstance() *CharServerInstance {
	if len(cs.Instances) == 0 {
		return nil
	}

	i := rand.Int() % len(cs.Instances)

	return cs.Instances[i]
}

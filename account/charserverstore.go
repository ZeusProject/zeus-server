package account

type CharServerStore interface {
	Servers() ([]*CharServer, error)
	RegisterInstance(id string, instance *CharServerInstance) error
	UpdateOnlineCount(id string, count int) error
	DeregisterInstance(id string, instance *CharServerInstance) error
}

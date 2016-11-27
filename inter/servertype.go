package inter

type ServerType int

const (
	_ ServerType = iota

	AccountServer
	CharServer
	ZoneServer
)

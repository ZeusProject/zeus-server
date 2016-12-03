package net

type RpcHandler interface {
	OnDisconnect(err error)
}

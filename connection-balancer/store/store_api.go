package store

import "github.com/hemantsharma1498/rtc/server/types"

type Storage interface {
	GetAllCommunicationServers() ([]*types.CommunicationServer, error)
}

type Connecter interface {
	Connect() (Storage, error)
}

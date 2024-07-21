package store

import "github.com/hemantsharma1498/rtc/server/types"



type Storage interface {
	SaveMessage(message *types.Message) error
	GetMessages(channelIds []int) ([]*types.Message, error)
}

type Connecter interface {
	Connect() (Storage, error)
}

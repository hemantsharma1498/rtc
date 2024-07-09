package store

import (
	"context"
	"members/store/types"
)

type Storage interface {
	CreateAccount(context.Context, *types.MemberAccount) (*types.MemberAccount, error)
	GetMembersByEmail(context.Context, []string) ([]*types.MemberAccount, error)
}

type Connecter interface {
	Connect() (Storage, error)
}

package users

import (
	"context"
	"github.com/go-rel/rel"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, user *User) error {
	return d.repository.Delete(ctx, user)
}

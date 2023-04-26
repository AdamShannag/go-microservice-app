package users

import (
	"context"
	"github.com/go-rel/rel"
)

type get struct {
	repository rel.Repository
}

func (g get) Get(ctx context.Context, users *User, id string) error {
	return g.repository.Find(ctx, users, rel.Select().Where(rel.Eq("id", id)))
}

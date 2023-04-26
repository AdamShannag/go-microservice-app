package users

import (
	"context"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, user *User) error {
	if err := user.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	return c.repository.Insert(ctx, user)
}

package users

import (
	"context"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, user *User, changes rel.Changeset) error {
	if err := user.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	return u.repository.Update(ctx, user, changes)
}

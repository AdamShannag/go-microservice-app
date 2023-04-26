package users

import (
	"context"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "user")))
)

//go:generate mockery --name=Service --case=underscore --output userstest --outpkg userstest

// Service instance for user's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	Get(ctx context.Context, users *User, id string) error
	GetAll(ctx context.Context, users *[]User, filter Filter) error
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User, changes rel.Changeset) error
	Delete(ctx context.Context, user *User)
}

// beside embeding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	get
	getAll
	create
	update
	delete
}

var _ Service = (*service)(nil)

// New User service.
func New(repository rel.Repository) Service {
	return service{
		get:    get{repository: repository},
		getAll: getAll{repository: repository},
		create: create{repository: repository},
		update: update{repository: repository},
		delete: delete{repository: repository},
	}
}

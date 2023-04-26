package users

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/reltest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdate_Update(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		id         = uuid.New()
		user       = User{
			ID:        id.String(),
			FirstName: "mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
		changes = rel.NewChangeset(&user)
	)

	user.FirstName = "Adam"

	repository.ExpectUpdate(changes).ForType("users.User")

	assert.Nil(t, service.Update(ctx, &user, changes))
	assert.NotEmpty(t, user.ID)

	repository.AssertExpectations(t)
}

func TestUpdate_validateError(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		id         = uuid.New()
		user       = User{
			ID:       id.String(),
			LastName: "Shnq",
			Email:    "adam@shnq.com",
			Phone:    "0798099158",
		}
		changes = rel.NewChangeset(&user)
	)

	assert.Equal(t, ErrUserFirstNameBlank, service.Update(ctx, &user, changes))

	repository.AssertExpectations(t)
}

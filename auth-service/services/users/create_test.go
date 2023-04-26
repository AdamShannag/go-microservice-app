package users

import (
	"context"
	"github.com/go-rel/reltest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate_Create(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		user       = User{
			ID:        uuid.New().String(),
			FirstName: "Mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
	)

	repository.ExpectInsert().For(&user)
	assert.Nil(t, service.Create(ctx, &user))
	assert.NotEmpty(t, user.ID)
	repository.AssertExpectations(t)
}

func TestCreate_validateError(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		user       = User{
			ID:       uuid.New().String(),
			LastName: "Shnq",
			Email:    "adam@shnq.com",
			Phone:    "0798099158",
		}
	)

	assert.Equal(t, ErrUserFirstNameBlank, service.Create(ctx, &user))
	repository.AssertExpectations(t)
}

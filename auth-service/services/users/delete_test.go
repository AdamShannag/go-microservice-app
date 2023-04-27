package users

import (
	"context"
	"github.com/go-rel/reltest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete_Delete(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		user       = User{
			ID:        uuid.New(),
			FirstName: "Mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
	)

	repository.ExpectDelete().ForType("users.User")

	assert.NotPanics(t, func() {
		service.Delete(ctx, &user)
	})

	repository.AssertExpectations(t)
}

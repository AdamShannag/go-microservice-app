package users

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/reltest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet_Get(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		user       User
		id         = uuid.New()
		result     = User{
			ID:        id,
			FirstName: "mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
	)

	repository.ExpectFind(
		rel.Select().Where(rel.Eq("id", id.String())),
	).Result(result)

	assert.NotPanics(t, func() {
		service.Get(ctx, &user, id.String())
		assert.Equal(t, result, user)
	})

	repository.AssertExpectations(t)
}

package users

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll_GetAll(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		filter     = Filter{"mo"}
		result     = []User{
			{
				ID:        uuid.New().String(),
				FirstName: "mohammad",
				LastName:  "Shnq",
				Email:     "adam@shnq.com",
				Phone:     "0798099158",
			},
		}
		users []User
	)

	repository.ExpectFindAll(
		rel.Where(where.Or(where.Like("first_name", "%mo%"), where.Like("last_name", "%mo%"))).SortAsc("first_name"),
	).Result(result)

	assert.NotPanics(t, func() {
		service.GetAll(ctx, &users, filter)
		assert.Equal(t, result, users)
	})

	repository.AssertExpectations(t)
}

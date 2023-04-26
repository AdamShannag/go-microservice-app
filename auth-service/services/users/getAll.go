package users

import (
	"context"
	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
	Name string
}

type getAll struct {
	repository rel.Repository
}

func (g getAll) GetAll(ctx context.Context, users *[]User, filter Filter) error {
	var (
		query = rel.Select().SortAsc("first_name")
	)

	if filter.Name != "" {
		query = query.
			Where(rel.Like("first_name", "%"+filter.Name+"%")).
			OrWhere(rel.Like("last_name", "%"+filter.Name+"%"))
	}

	return g.repository.FindAll(ctx, users, query)
}

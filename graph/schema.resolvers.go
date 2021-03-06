package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/asSqr/graphql-back/graph/generated"
	"github.com/asSqr/graphql-back/graph/model"
)

func (r *queryResolver) Recipe(ctx context.Context, id int) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

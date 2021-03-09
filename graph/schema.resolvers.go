package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"net/http"

	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateThought(ctx context.Context, input model.NewThought) (*model.Thought, error) {
	res := r.srv.CreateNew(input)
	if res == nil {
		return nil, gqlerror.Errorf("(%d) Internal Server Error", http.StatusInternalServerError)
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) UpdateThought(ctx context.Context, id int, userID int, input model.NewThought) (*model.Thought, error) {
	res, err := r.srv.UpdateOne(id, userID, input)
	if err != nil {
		return nil, gqlerror.Errorf("(%d) %s", err.Type, err.Response)
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) DeleteThought(ctx context.Context, id int, userID int) (*model.Thought, error) {
	res, err := r.srv.DeleteOne(id, userID)
	if err != nil {
		return nil, gqlerror.Errorf("(%d) %s", err.Type, err.Response)
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) Thoughts(ctx context.Context) ([]*model.Thought, error) {
	return r.srv.GetAll().ToGraphQLs(), nil
}

func (r *queryResolver) Thought(ctx context.Context, id int) (*model.Thought, error) {
	return r.srv.GetOne(id).ToGraphQL(), nil
}

func (r *thoughtResolver) User(ctx context.Context, obj *model.Thought) (*model.User, error) {
	return &model.User{
		ID:    obj.UserID,
		Name:  "anom",
		Email: "someone-somewhere@gmail.com",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Thought returns generated.ThoughtResolver implementation.
func (r *Resolver) Thought() generated.ThoughtResolver { return &thoughtResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type thoughtResolver struct{ *Resolver }

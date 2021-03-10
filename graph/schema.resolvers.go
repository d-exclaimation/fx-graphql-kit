package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
	"github.com/d-exclaimation/fx-graphql-kit/server/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	res, err := r.usrv.NewUser(name, email)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) CreateThought(ctx context.Context, input model.NewThought) (*model.Thought, error) {
	res, err := r.srv.CreateNew(input)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) UpdateThought(ctx context.Context, id int, userID int, input model.NewThought) (*model.Thought, error) {
	res, err := r.srv.UpdateOne(id, userID, input)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) DeleteThought(ctx context.Context, id int, userID int) (*model.Thought, error) {
	res, err := r.srv.DeleteOne(id, userID)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) Thoughts(ctx context.Context) ([]*model.Thought, error) {
	res, err := r.srv.GetAll()
	if err != nil {
		return nil, err.ToGQLError()
	}

	return res.ToGraphQLs(), nil
}

func (r *queryResolver) Thought(ctx context.Context, id int) (*model.Thought, error) {
	res, err := r.srv.GetOne(id)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *thoughtResolver) User(ctx context.Context, obj *model.Thought) (*model.User, error) {
	// ID is usually string for like UUID but in this context I used uint
	id, err := strconv.ParseUint(obj.UserID, 10, 64)
	if err != nil {
		return nil, gqlerror.Errorf("Cannot find User")
	}

	// Get the data from dataloader given inside the context object
	user, err := middleware.For(ctx).UserById.Load(uint(id))
	if err != nil || user == nil {
		return nil, err
	}
	return user.ToGraphQL(), nil
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

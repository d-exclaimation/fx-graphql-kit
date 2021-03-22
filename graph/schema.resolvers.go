package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
)

func (r *mutationResolver) CreateThought(ctx context.Context, input model.NewThought) (*model.Thought, error) {
	res, err := r.srv.CreateNew(ctx, input)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) UpdateThought(ctx context.Context, id int, userID int, input model.NewThought) (*model.Thought, error) {
	res, err := r.srv.UpdateOne(ctx, id, userID, input)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) DeleteThought(ctx context.Context, id int, userID int) (*model.Thought, error) {
	res, err := r.srv.DeleteOne(ctx, id, userID)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) Thoughts(ctx context.Context) ([]*model.Thought, error) {
	res, err := r.srv.GetAll(ctx)
	if err != nil {
		return nil, err.ToGQLError()
	}

	return res.ToGraphQLs(), nil
}

func (r *queryResolver) Thought(ctx context.Context, id int) (*model.Thought, error) {
	res, err := r.srv.GetOne(ctx, id)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return res.ToGraphQL(), nil
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

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
)

func (r *mutationResolver) CreateThought(ctx context.Context, input model.NewThought) (*model.Thought, error) {
	newThought := &model.Thought{
		ID:       "1",
		Title:    input.Title,
		Body:     input.Body,
		ImageURL: input.ImageURL,
	}
	return newThought, nil
}

func (r *mutationResolver) UpdateThought(ctx context.Context, id int, userID int, input *model.NewThought) (*model.Thought, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteThought(ctx context.Context, id int, userID int) (*model.Thought, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Thoughts(ctx context.Context) ([]*model.Thought, error) {
	return make([]*model.Thought, 0), nil
}

func (r *queryResolver) Thought(ctx context.Context, id int) (*model.Thought, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *thoughtResolver) User(ctx context.Context, obj *model.Thought) (*model.User, error) {
	return &model.User{
		ID:    obj.UserID,
		Name:  "anom",
		Email: "vincentlimchen@gmail.com",
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

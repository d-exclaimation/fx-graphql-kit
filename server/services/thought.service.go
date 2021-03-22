//
//  thought.service.go
//  services
//
//  Created by d-exclaimation on 8:17 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"context"
	"github.com/d-exclaimation/fx-graphql-kit/ent"
	"github.com/d-exclaimation/fx-graphql-kit/ent/thought"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
	"github.com/d-exclaimation/fx-graphql-kit/server/errors"
	"net/http"
)

// ThoughtService Struct
type ThoughtService struct {
	client *ent.Client
}

// Fx Provider
func ThoughtServiceProvider(db *ent.Client) *ThoughtService {
	return &ThoughtService{
		client: db,
	}
}

// Methods
func (srv *ThoughtService) CreateNew(ctx context.Context, input model.NewThought) (*ent.Thought, *errors.ServiceError) {
	// Grab the client and create and fill in all fields
	res, err := srv.client.
		Thought.Create().
		SetTitle(input.Title).
		SetBody(input.Body).
		SetImageURL(handleImageURL(input.ImageURL)).
		SetUserId(int64(input.UserID)).
		Save(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (srv *ThoughtService) GetAll(ctx context.Context) (ent.ThoughtsArray, *errors.ServiceError) {
	res, err := srv.client.
		Thought.Query().
		All(ctx)
	if err != nil {
		return make(ent.ThoughtsArray, 0), errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (srv *ThoughtService) GetOne(ctx context.Context, id int) (*ent.Thought, *errors.ServiceError) {
	res, err := srv.client.
		Thought.Query().
		Where(thought.ID(id)).
		First(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (srv *ThoughtService) UpdateOne(ctx context.Context, id int, userId int, input model.NewThought) (*ent.Thought, *errors.ServiceError) {
	curr, fail := srv.GetOne(ctx, id)
	if fail != nil {
		return nil, fail
	}

	if curr.UserId != int64(userId) {
		return nil, errors.NewServiceError(http.StatusForbidden, "Unauthorized permission to update data")
	}

	res, err := srv.client.
		Thought.UpdateOneID(id).
		SetTitle(input.Title).
		SetBody(input.Body).
		SetImageURL(handleImageURL(input.ImageURL)).
		Save(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (srv *ThoughtService) DeleteOne(ctx context.Context, id int, userId int) (*ent.Thought, *errors.ServiceError) {
	curr, fail := srv.GetOne(ctx, id)
	if fail != nil {
		return nil, fail
	}

	if curr.UserId != int64(userId) {
		return nil, errors.NewServiceError(http.StatusForbidden, "Unauthorized permission to update data")
	}

	if err := srv.client.Thought.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return curr, nil
}


func handleImageURL(url *string) string {
	// For some reason ent and gqlgen uses different technique for nullable
	imageurl := ""
	if url != nil {
		imageurl = *url
	}
	return imageurl
}
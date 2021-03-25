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
func (srv *ThoughtService) CreateNew(ctx context.Context, input model.NewThought) (*ent.Thought, error) {
    // Query: INSERT INTO thoughts (title, body, image_url, user_id) VALUES (?, ?, ?, ?) RETURNING *;
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

func (srv *ThoughtService) GetAll(ctx context.Context) (ent.ThoughtsArray, error) {
    // Query: SELECT * FROM thoughts;
	res, err := srv.client.
		Thought.Query().
		All(ctx)
	if err != nil {
		return make(ent.ThoughtsArray, 0), errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (srv *ThoughtService) GetOne(ctx context.Context, id int) (*ent.Thought, error) {
    // Query: SELECT * FROM thoughts WHERE id=?;
	res, err := srv.client.
		Thought.Query().
		Where(thought.ID(id)).
		First(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (srv *ThoughtService) UpdateOne(ctx context.Context, id int, userId int, input model.NewThought) (*ent.Thought, error) {
	curr, err := srv.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if curr.UserId != int64(userId) {
		return nil, errors.NewServiceError(http.StatusForbidden, "Unauthorized permission to update data")
	}

    // Query: UPDATE thoughts SET title=?, body=?, image_url=? WHERE id=? RETURNING *;
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

func (srv *ThoughtService) DeleteOne(ctx context.Context, id int, userId int) (*ent.Thought, error) {
	curr, err := srv.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if curr.UserId != int64(userId) {
		return nil, errors.NewServiceError(http.StatusForbidden, "Unauthorized permission to update data")
	}

    // Query: DELETE FROM thoughts WHERE id=?;
	if err := srv.client.Thought.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return curr, nil
}


func handleImageURL(url *string) string {
	imageurl := ""
	if url != nil {
		imageurl = *url
	}
	return imageurl
}

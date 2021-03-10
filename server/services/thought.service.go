//
//  thought.service.go
//  services
//
//  Created by d-exclaimation on 8:17 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"github.com/d-exclaimation/fx-graphql-kit/db/entities"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
	"github.com/d-exclaimation/fx-graphql-kit/server/errors"
	"gorm.io/gorm"
	"net/http"
)

// ThoughtService Struct
type ThoughtService struct {
	usrv *UserService
	db *gorm.DB
}

// Fx Provider
func ThoughtServiceProvider(db *gorm.DB, usrv *UserService) *ThoughtService {
	return &ThoughtService{
		db: db,
		usrv: usrv,
	}
}

// Methods
func (srv *ThoughtService) CreateNew(input model.NewThought) (*entities.Thought, *errors.ServiceError) {
	uid := uint(input.UserID)
	thought := &entities.Thought{
		Title:    input.Title,
		Body:     input.Body,
		ImageURL: input.ImageURL,
		UserID:   uint(input.UserID),
	}

	if err := srv.usrv.AppendThought(uid, thought); err != nil {
		return nil, err
	}
	return thought, nil
}

func (srv *ThoughtService) GetAll() (entities.ThoughtsArray, *errors.ServiceError) {
	var todos entities.ThoughtsArray
	if err := srv.db.Find(&todos).Error; err != nil {
		return make([]*entities.Thought, 0), errors.NewServiceError(http.StatusInternalServerError, "Cannot fetch data from database")
	}
	return todos, nil
}

func (srv *ThoughtService) GetOne(id int) (*entities.Thought, *errors.ServiceError) {
	thoughts, err := srv.GetAll()
	if err != nil {
		return nil, err
	}

	for _, thought := range thoughts {
		if thought.ID == uint(id) {
			return thought, nil
		}
	}
	return nil, errors.NewServiceError(http.StatusNotFound, "Cannot find Thought, Invalid ID")
}

func (srv *ThoughtService) GetOwner(id int) (*entities.User, *errors.ServiceError) {
	curr, fail := srv.GetOne(id)
	if fail != nil {
		return nil, fail
	}

	user, fail := srv.usrv.GetUser(curr.UserID)
	if fail != nil {
		return nil, fail
	}
	return user, nil
}


func (srv *ThoughtService) UpdateOne(id int, userId int, input model.NewThought) (*entities.Thought, *errors.ServiceError) {
	selected, err := srv.GetOne(id)

	// Errors
	if err != nil {
		return nil, err
	}
	if selected.UserID != uint(userId) {
		return nil, errors.NewServiceError(http.StatusForbidden, "Invalid Permission")
	}

	// Retrieve, Update, and Save
	srv.db.First(selected)

	selected.Title = input.Title
	selected.Body = input.Body
	selected.ImageURL = input.ImageURL

	srv.db.Save(selected)

	return selected, nil
}

func (srv *ThoughtService) DeleteOne(id int, userId int) (*entities.Thought, *errors.ServiceError) {
	selected, err := srv.GetOne(id)
	if err != nil  {
		return nil, err
	}

	if selected.UserID != uint(userId) {
		return nil, errors.NewServiceError(http.StatusForbidden, "Invalid Permission")
	}

	copied := &entities.Thought{
		Model:    selected.Model,
		Title:    selected.Title,
		Body:     selected.Body,
		ImageURL: selected.ImageURL,
		UserID:   selected.UserID,
	}

	srv.db.Delete(selected)
	return copied, nil
}
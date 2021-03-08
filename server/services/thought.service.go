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
	"gorm.io/gorm"
)

type ThoughtService struct {
	db *gorm.DB
}

func ThoughtServiceProvider(db *gorm.DB) *ThoughtService {
	return &ThoughtService{
		db: db,
	}
}

func (srv *ThoughtService) CreateNew(input model.NewThought) *entities.Thought {
	thought := &entities.Thought{
		Title:    input.Title,
		Body:     input.Body,
		ImageURL: input.ImageURL,
		UserID:   uint(input.UserID),
	}
	if err := srv.db.Create(thought).Error; err != nil {
		return nil
	}
	return thought
}

func (srv *ThoughtService) GetAll() entities.ThoughtsArray {
	var todos entities.ThoughtsArray
	if err := srv.db.Find(&todos).Error; err != nil {
		return make([]*entities.Thought, 0)
	}
	return todos
}

func (srv *ThoughtService) GetOne(id int) *entities.Thought {
	thoughts := srv.GetAll()
	for _, thought := range thoughts {
		if thought.ID == uint(id) {
			return thought
		}
	}
	return nil
}

func (srv *ThoughtService) UpdateOne(id int, userId int, input model.NewThought) *entities.Thought {
	selected := srv.GetOne(id)
	if selected == nil || selected.UserID != uint(userId) {
		return nil
	}

	// Retrieve, Update, and Save
	srv.db.First(selected)

	selected.Title = input.Title
	selected.Body = input.Body
	selected.ImageURL = input.ImageURL

	srv.db.Save(selected)

	return selected
}

func (srv *ThoughtService) DeleteOne(id int, userId int) *entities.Thought {
	selected := srv.GetOne(id)
	if selected == nil || selected.UserID != uint(userId) {
		return nil
	}

	copied := &entities.Thought{
		Model:    selected.Model,
		Title:    selected.Title,
		Body:     selected.Body,
		ImageURL: selected.ImageURL,
		UserID:   selected.UserID,
	}

	srv.db.Delete(selected)
	return copied
}
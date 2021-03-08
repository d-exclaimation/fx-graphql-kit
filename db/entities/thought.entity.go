//
//  thought.entities.go
//  db
//
//  Created by d-exclaimation on 7:24 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package entities

import (
	"fmt"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
	"gorm.io/gorm"
)

// Thought Database Entity
type Thought struct {
	gorm.Model
	Title     string
	Body      string
	ImageURL  *string
	UserID 	  uint
}

// Convert to GraphQL Schema
func (s *Thought) ToGraphQL() *model.Thought {
	return &model.Thought{
		ID:       fmt.Sprintf("%d", s.ID),
		Title:    s.Title,
		Body:     s.Body,
		ImageURL: s.ImageURL,
		UserID:   fmt.Sprintf("%d", s.UserID),
	}
}

// Method Injection
type ThoughtsArray []*Thought

// Convert all to GraphQL Schema
func (arr ThoughtsArray) ToGraphQLs() []*model.Thought {
	res := make([]*model.Thought, len(arr))
	for i, thought := range arr {
		res[i] = thought.ToGraphQL()
	}
	return res
}
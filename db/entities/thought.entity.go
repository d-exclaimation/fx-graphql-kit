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


type Thought struct {
	gorm.Model
	Title     string
	Body      string
	ImageURL  *string
}

func (s *Thought) ToGraphQL() *model.Thought {
	return &model.Thought{
		ID:       fmt.Sprintf("%d", s.ID),
		Title:    s.Title,
		Body:     s.Body,
		ImageURL: s.ImageURL,
	}
}

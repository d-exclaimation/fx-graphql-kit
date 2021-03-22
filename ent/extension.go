//
//  extension.go
//  extensions
//
//  Created by d-exclaimation on 10:09 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package ent

import (
	"fmt"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
)

// Convert to GraphQL Schema
func (t *Thought) ToGraphQL() *model.Thought {
	return &model.Thought{
		ID:       fmt.Sprintf("%d", t.ID),
		Title:    t.Title,
		Body:     t.Body,
		ImageURL: &t.ImageURL,
		UserID:   fmt.Sprintf("%d", t.UserId),
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

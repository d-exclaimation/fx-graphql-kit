//
//  user.entity.go
//  entities
//
//  Created by d-exclaimation on 6:15 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package entities

import (
	"fmt"
	"github.com/d-exclaimation/fx-graphql-kit/graph/model"
	"gorm.io/gorm"
)

// User Database Entity
type User struct {
	gorm.Model
	Name      string
	Email 	  string
	Thoughts  []Thought `gorm:"foreignKey:UserID"`
}

func (user *User) ToGraphQL() *model.User {
	return &model.User{
		ID:    fmt.Sprintf("%d", user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}
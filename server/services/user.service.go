//
//  user.service.go
//  services
//
//  Created by d-exclaimation on 6:31 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"github.com/d-exclaimation/fx-graphql-kit/db/entities"
	"github.com/d-exclaimation/fx-graphql-kit/server/errors"
	"gorm.io/gorm"
	"net/http"
)

// UserService Struct
type UserService struct {
	db *gorm.DB
}

// Fx Provider
func UserServiceProvider(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (srv *UserService) NewUser(name string, email string) (*entities.User, *errors.ServiceError) {
	user := &entities.User{
		Name:     name,
		Email:    email,
		Thoughts: make([]entities.Thought, 0),
	}

	if err := srv.db.Create(user).Error; err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, "Fail to create User")
	}

	return user, nil
}

func (srv *UserService) GetSome(uids []uint) ([]*entities.User, error) {
	// Resolver for multiple user
	var users []*entities.User
	if err := srv.db.Find(&users).Error; err != nil {
		return make([]*entities.User, 0), err
	}

	// Make a hashmap to reduce time complexity
	usersMap := make(map[uint]*entities.User)
	for _, user := range users {
		usersMap[user.ID] = user
	}

	// Find by hashmap
	res := make([]*entities.User, len(uids))
	for i, uid := range uids {
		res[i] = usersMap[uid]
	}

	return users, nil
}

func (srv *UserService) GetUser(uid uint) (*entities.User, *errors.ServiceError) {
	user := &entities.User{Model: gorm.Model{ID: uid}}
	if err := srv.db.First(user).Error; err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, "Fail to find User")
	}
	return user, nil
}

func (srv *UserService) AppendThought(uid uint, thought *entities.Thought) *errors.ServiceError {
	// Get the user database model and append a item aka Create with relation
	if err := srv.db.Model(&entities.User{
		Model: gorm.Model{
			ID: uid,
		},
	}).Association("Thoughts").Append(thought); err != nil {
		return errors.NewServiceError(http.StatusInternalServerError, "Fail to create User")
	}
	return nil
}
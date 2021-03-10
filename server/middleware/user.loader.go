//
//  user.loader.go
//  data
//
//  Created by d-exclaimation on 7:40 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"context"
	"github.com/d-exclaimation/fx-graphql-kit/db/entities"
	"github.com/d-exclaimation/fx-graphql-kit/server/services"
	"github.com/gin-gonic/gin"
	"time"
)

const loadersKey = "userloaders"

type Loaders struct {
	UserById *entities.UserLoader
}

func UserLoaderMiddleWareProvider(usrv *services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newContext := context.WithValue(ctx.Request.Context(), loadersKey, &Loaders{
			UserById: entities.NewUserLoader(entities.UserLoaderConfig{
				Fetch: func(keys []uint) ([]*entities.User, []error) {
					// Basically resolver for all concurrent request given the ids
					res, err := usrv.GetSome(keys)
					if err != nil {
						return nil, []error{err}
					}

					return res, nil
				},
				Wait:     1 * time.Millisecond,
				MaxBatch: 100,
			}),
		})
		ctx.Request = ctx.Request.WithContext(newContext)
		ctx.Next()
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
	/*
	Gin context does not work, use regular context
	Data loading
	*/
}
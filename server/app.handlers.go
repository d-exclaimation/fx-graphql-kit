//
//  app.handlers.go
//  server
//
//  Created by d-exclaimation on 8:11 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppHandlers struct {
	Middlewares []gin.HandlerFunc
	GQLHandler  gin.HandlerFunc
	Playground  gin.HandlerFunc
}

func AppHandlersProvider(module generated.Config) *AppHandlers {
	return &AppHandlers{
		Middlewares: []gin.HandlerFunc{middleware.GinContextToContextMiddleware()},
		GQLHandler:  GraphqlHandler(module),
		Playground:  PlaygroundHandler(),
	}
}

func GraphqlHandler(module generated.Config) gin.HandlerFunc {
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(module))
	return func(ctx *gin.Context) {
		graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	playgroundHandler := playground.Handler("Nodes-Graph API Playground", graphqlPath)
	return func(ctx *gin.Context) {
		playgroundHandler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

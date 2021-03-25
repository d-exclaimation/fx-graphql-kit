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
	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

// AppHandlers / Controller
type AppHandlers struct {
	Middlewares []echo.MiddlewareFunc
	GQLHandler  echo.HandlerFunc
	Playground  echo.HandlerFunc
}

// Fx Provider
func AppHandlersProvider(module generated.Config) *AppHandlers {
	return &AppHandlers{
		Middlewares: []echo.MiddlewareFunc{
			middleware.EchoContextMiddleware,
            // TODO: Move to middleware directory
			em.LoggerWithConfig(em.LoggerConfig{
				Format: "${time_rfc3339_nano} |" + color.NewRGBStyle(color.RGB(200, 200, 200), color.HEX("#20bcaf", true)).Sprint(" ${status} ${method} ") + "| ${latency_human} | >> ${uri}\n",
			}),
		},
		GQLHandler:  GraphqlHandler(module),
		Playground:  PlaygroundHandler(),
	}
}

// GraphQL Query Handler
func GraphqlHandler(module generated.Config) echo.HandlerFunc {
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(module))
	return func(ctx echo.Context) error {
		graphqlServer.ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// Playground Handler
func PlaygroundHandler() echo.HandlerFunc {
	playgroundHandler := playground.Handler("Nodes-Graph API Playground", graphqlPath)
	return func(ctx echo.Context) error {
		playgroundHandler.ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

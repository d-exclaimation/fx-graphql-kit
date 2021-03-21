//
//  app.go
//  server
//
//  Created by d-exclaimation on 8:05 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package server

import (
	"context"
	"github.com/d-exclaimation/fx-graphql-kit/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const (
	graphqlPath = "/graphql"
	entry = "/"
)

// Fx Provider
func AppProvider(lifecycle fx.Lifecycle) *echo.Echo {
	app := echo.New()
	port := config.GetPort()

	// Using Fx Lifecycle create start and stop functions to be invoke at appropriate condition
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go (func() {
				_ = app.Start(":" + port)
			})()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown(ctx)
		},
	})

	return app
}

// Fx Invoke Middleware
func InvokeMiddleWare(app *echo.Echo, handlers *AppHandlers) {
	for _, mw := range handlers.Middlewares {
		app.Use(mw)
	}
}

// Fx Invoke Handler
func InvokeHandler(app *echo.Echo, handlers *AppHandlers) {
	app.POST(graphqlPath, handlers.GQLHandler)
	app.GET(entry, handlers.Playground)
}
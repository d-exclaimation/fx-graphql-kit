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
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

const (
	graphqlPath = "/graphql"
	entry = "/"
)

func AppProvider(lifecycle fx.Lifecycle) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	port := config.GetPort()
	
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           app,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go (func() {
				_ = srv.ListenAndServe()
			})()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return app
}

func InvokeMiddleWare(app *gin.Engine, handlers *AppHandlers) {
	for _, mw := range handlers.Middlewares {
		app.Use(mw)
	}
}

func InvokeHandler(app *gin.Engine, handlers *AppHandlers) {
	app.POST(graphqlPath, handlers.GQLHandler)
	app.GET(entry, handlers.Playground)
}
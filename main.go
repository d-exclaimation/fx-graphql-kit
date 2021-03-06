//
//  schema.resolvers.go
//  fx-graphql-kit
//
//  Created by d-exclaimation on 8:24 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package main

import (
	"github.com/d-exclaimation/fx-graphql-kit/db"
	"github.com/d-exclaimation/fx-graphql-kit/graph"
	"github.com/d-exclaimation/fx-graphql-kit/server"
	"github.com/d-exclaimation/fx-graphql-kit/server/services"
	"go.uber.org/fx"
)

// Fx Runtime Lifecycle
func main() {
	fx.New(
		fx.Provide(
			// Server HTTP App
			server.AppProvider,

			// Database connection / client
			db.EntProvider,

			// Services and Modules / Configuration
			services.ThoughtServiceProvider,
			graph.ModuleProvider,

			// Handlers / Controllers
			server.AppHandlersProvider,
		),
		fx.Invoke(
			// HTTP Middleware and Endpoints Invoker
			server.InvokeMiddleWare,
			server.InvokeHandler,
		),
	).Run()
}

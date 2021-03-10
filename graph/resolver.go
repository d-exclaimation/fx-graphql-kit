package graph

import (
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/server/services"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver Struct
type Resolver struct {
	srv *services.ThoughtService
	usrv *services.UserService
}

// Resolver Constructor
func NewResolver(srv *services.ThoughtService, usrv *services.UserService) *Resolver {
	return &Resolver{
		srv: srv,
		usrv: usrv,
	}
}

// Fx Provider
func ModuleProvider(srv *services.ThoughtService, usrv *services.UserService) generated.Config {
	return generated.Config {
		Resolvers: NewResolver(srv, usrv),
	}
}
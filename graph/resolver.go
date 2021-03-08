package graph

import (
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/server/services"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver Struct
type Resolver struct {
	srv *services.ThoughtService
}

// Resolver Constructor
func NewResolver(srv *services.ThoughtService) *Resolver {
	return &Resolver{
		srv: srv,
	}
}

// Fx Provider
func ModuleProvider(srv *services.ThoughtService) generated.Config {
	return generated.Config {
		Resolvers: NewResolver(srv),
	}
}
package graph

import (
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"github.com/d-exclaimation/fx-graphql-kit/server/services"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	srv *services.ThoughtService
}

func NewResolver(srv *services.ThoughtService) *Resolver {
	return &Resolver{
		srv: srv,
	}
}

func ModuleProvider(srv *services.ThoughtService) generated.Config {
	return generated.Config {
		Resolvers: NewResolver(srv),
	}
}
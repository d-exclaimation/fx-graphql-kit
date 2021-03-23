package graph

import (
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver Struct
type Resolver struct {
}

// Resolver Constructor
func NewResolver() *Resolver {
	return &Resolver{
	}
}

// Fx Provider
func ModuleProvider() generated.Config {
	return generated.Config {
		Resolvers: NewResolver(),
	}
}
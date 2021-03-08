package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/d-exclaimation/fx-graphql-kit/config"
	"github.com/d-exclaimation/fx-graphql-kit/graph"
	"github.com/d-exclaimation/fx-graphql-kit/graph/generated"
	"log"
	"net/http"
)


func main() {
	port := config.GetPort()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

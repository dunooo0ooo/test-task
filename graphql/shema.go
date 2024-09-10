package graphql

import (
	"github.com/99designs/gqlgen/graphql/complexity"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
)

func NewGraphQLServer(resolver *Resolver) *handler.Server {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))

	srv.Use(&lru.Cache{MaxSize: 100})
	srv.Use(complexity.Limiter{ComplexityLimit: 100})

	return srv
}

func NewPlaygroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/query")
}

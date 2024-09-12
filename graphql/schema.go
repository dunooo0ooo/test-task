package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
)

func NewGraphQLServer(resolver *Resolver) *handler.Server {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))
	return srv
}

func NewPlaygroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/query")
}

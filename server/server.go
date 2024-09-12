package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"test-task/graphql"
)

func main() {
	posts := make(map[string]*graphql.Post)
	comments := make(map[string]*graphql.Comment)

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{
			PostsMap:    posts,
			CommentsMap: comments,
		},
	}))

	http.Handle("/query", srv)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	log.Println("Server is running on http://localhost:8080/playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

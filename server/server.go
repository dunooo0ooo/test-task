package main

import (
	"log"
	"net/http"
	"test-task/graphql"
	"test-task/repository"
	"test-task/services"
)

func main() {
	repo := repository.NewInMemoryRepository()

	postService := services.NewPostService(repo)
	commentService := services.NewCommentService(repo)

	resolver := graphql.NewResolver(postService, commentService)

	srv := graphql.NewGraphQLServer(resolver)

	http.Handle("/query", srv)
	http.Handle("/playground", graphql.NewPlaygroundHandler())

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

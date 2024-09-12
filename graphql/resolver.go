package graphql

import "test-task/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	postService    *services.PostService
	commentService *services.CommentService
}

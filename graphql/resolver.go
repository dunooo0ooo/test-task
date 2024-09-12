package graphql

import "test-task/services"

type Resolver struct {
	postService    *services.PostService
	commentService *services.CommentService
	PostsMap       map[string]*Post
	CommentsMap    map[string]*Comment
}

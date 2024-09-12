package graphql

import (
	"context"
	"test-task/models"
	"test-task/services"
)

func NewResolver(postService *services.PostService, commentService *services.CommentService) *Resolver {
	return &Resolver{
		postService:    postService,
		commentService: commentService,
	}
}

func (r *Resolver) Query_posts(ctx context.Context) ([]*models.Post, error) {
	return r.postService.GetPosts(), nil
}

func (r *Resolver) Query_post(ctx context.Context, id string) (*models.Post, error) {
	return r.postService.GetPostByID(id)
}

func (r *Resolver) Mutation_createPost(ctx context.Context, title string, content string) (*models.Post, error) {
	return r.postService.CreatePost(title, content), nil
}

func (r *Resolver) Mutation_createComment(ctx context.Context, postId string, content string, parentId *string) (*models.Comment, error) {
	return r.commentService.CreateComment(postId, content, parentId)
}

func (r *Resolver) Mutation_lockPost(ctx context.Context, postId string) (*models.Post, error) {
	err := r.postService.LockPost(postId)
	if err != nil {
		return nil, err
	}
	return r.postService.GetPostByID(postId)
}

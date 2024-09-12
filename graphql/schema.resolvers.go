package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string) (*Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, postID string, content string, parentID *string) (*Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// LockPost is the resolver for the lockPost field.
func (r *mutationResolver) LockPost(ctx context.Context, postID string) (*Post, error) {
	panic(fmt.Errorf("not implemented: LockPost - lockPost"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

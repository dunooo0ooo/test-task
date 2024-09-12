package graphql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string) (*Post, error) {
	id := uuid.New().String()

	post := &Post{
		ID:       id,
		Title:    title,
		Content:  content,
		IsLocked: false,
		Comments: []*Comment{},
	}

	r.PostsMap[id] = post

	return post, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID string, content string, parentID *string) (*Comment, error) {
	post, exists := r.PostsMap[postID]
	if !exists {
		return nil, fmt.Errorf("post not found")
	}

	if post.IsLocked {
		return nil, fmt.Errorf("comments are disabled for this post")
	}

	id := uuid.New().String()

	comment := &Comment{
		ID:       id,
		PostID:   postID,
		Content:  content,
		ParentID: parentID,
		Replies:  []*Comment{},
	}

	if parentID != nil {
		parentComment, exists := r.CommentsMap[*parentID]
		if !exists {
			return nil, fmt.Errorf("parent comment not found")
		}
		parentComment.Replies = append(parentComment.Replies, comment)
	} else {
		post.Comments = append(post.Comments, comment)
	}

	r.CommentsMap[id] = comment

	return comment, nil
}

func (r *mutationResolver) LockPost(ctx context.Context, postID string) (*Post, error) {
	post, exists := r.PostsMap[postID]
	if !exists {
		return nil, fmt.Errorf("post not found")
	}

	post.IsLocked = true

	return post, nil
}

func (r *queryResolver) Posts(ctx context.Context, limit *int, page *int) ([]*Post, error) {
	lim := 10
	pg := 1
	if limit != nil {
		lim = *limit
	}
	if page != nil {
		pg = *page
	}

	start := (pg - 1) * lim
	end := start + lim

	if start >= len(r.PostsMap) {
		return []*Post{}, nil
	}

	if end > len(r.PostsMap) {
		end = len(r.PostsMap)
	}

	posts := []*Post{}
	i := 0
	for _, post := range r.PostsMap {
		if i >= start && i < end {
			posts = append(posts, post)
		}
		i++
	}

	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*Post, error) {
	post, exists := r.PostsMap[id]
	if !exists {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var posts = map[string]*Post{}
var comments = map[string]*Comment{}

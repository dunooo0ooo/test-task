package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"test-task/models"
)

type InMemoryRepository struct {
	posts    map[string]*models.Post
	comments map[string]*models.Comment
	mu       sync.Mutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		posts:    make(map[string]*models.Post),
		comments: make(map[string]*models.Comment),
	}
}

func (repo *InMemoryRepository) CreatePost(title, content string) *models.Post {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	post := &models.Post{
		ID:       uuid.New().String(),
		Title:    title,
		Content:  content,
		IsLocked: false,
		Comments: []*models.Comment{},
	}
	repo.posts[post.ID] = post
	return post
}

func (repo *InMemoryRepository) GetPosts() []*models.Post {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	posts := []*models.Post{}
	for _, post := range repo.posts {
		posts = append(posts, post)
	}
	return posts
}

func (repo *InMemoryRepository) GetPostByID(postID string) (*models.Post, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	post, exists := repo.posts[postID]
	if !exists {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (repo *InMemoryRepository) CreateComment(postID, content string, parentID *string) (*models.Comment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	post, exists := repo.posts[postID]
	if !exists {
		return nil, errors.New("post not found")
	}
	if post.IsLocked {
		return nil, errors.New("comments are locked for this post")
	}

	comment := &models.Comment{
		ID:       uuid.New().String(),
		PostID:   postID,
		ParentID: parentID,
		Content:  content,
		Replies:  []*models.Comment{},
	}

	if parentID == nil {
		post.Comments = append(post.Comments, comment)
	} else {
		parentComment, exists := repo.comments[*parentID]
		if !exists {
			return nil, errors.New("parent comment not found")
		}
		parentComment.Replies = append(parentComment.Replies, comment)
	}

	repo.comments[comment.ID] = comment
	return comment, nil
}

func (repo *InMemoryRepository) LockPost(postID string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	post, exists := repo.posts[postID]
	if !exists {
		return errors.New("post not found")
	}
	post.IsLocked = true
	return nil
}

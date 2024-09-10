package services

import (
	"test-task/models"
	"test-task/repository"
)

type PostService struct {
	repo *repository.InMemoryRepository
}

func NewPostService(repo *repository.InMemoryRepository) *PostService {
	return &PostService{repo: repo}
}

func (ps *PostService) CreatePost(title, content string) *models.Post {
	return ps.repo.CreatePost(title, content)
}

func (ps *PostService) GetPosts() []*models.Post {
	return ps.repo.GetPosts()
}

func (ps *PostService) GetPostByID(postID string) (*models.Post, error) {
	return ps.repo.GetPostByID(postID)
}

func (ps *PostService) LockPost(postID string) error {
	return ps.repo.LockPost(postID)
}

package services

import (
	"test-task/models"
	"test-task/repository"
)

type CommentService struct {
	repo *repository.InMemoryRepository
}

func NewCommentService(repo *repository.InMemoryRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (cs *CommentService) CreateComment(postID, content string, parentID *string) (*models.Comment, error) {
	return cs.repo.CreateComment(postID, content, parentID)
}

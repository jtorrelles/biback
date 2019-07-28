package services

import (
	"biback/app/models"
	"context"
)

// Usecase represent the article's usecases
type Showservice interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Show, string, error)
	GetByID(ctx context.Context, id int64) ([]*models.Show, error)
	Store(ctx context.Context, arg *models.Show) error
	Update(ctx context.Context, id int64, arg *models.Show) error
	//Update(ctx context.Context, ar *models.Article) error
	//GetByTitle(ctx context.Context, title string) (*models.Article, error)
	//Delete(ctx context.Context, id int64) error
}

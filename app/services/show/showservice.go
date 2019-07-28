package show

import (
	"biback/app/models"
	"biback/app/repository"
	"biback/app/services"
	"context"
	"time"
)

type showService struct {
	showRepo       repository.Repository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewShowService(a repository.Repository, timeout time.Duration) services.Showservice {
	return &showService{
		showRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *showService) Fetch(c context.Context, cursor string, num int64) ([]*models.Show, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listShow, nextCursor, err := a.showRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	/*listArticle, err = a.fillAuthorDetails(ctx, listArticle)
	if err != nil {
		return nil, "", err
	}*/

	return listShow, nextCursor, nil
}

func (a *showService) GetByID(c context.Context, id int64) ([]*models.Show, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listShow, err := a.showRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	/*listArticle, err = a.fillAuthorDetails(ctx, listArticle)
	if err != nil {
		return nil, "", err
	}*/

	return listShow, nil
}

func (a *showService) Store(c context.Context, m *models.Show) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	//Validar que no exista

	err := a.showRepo.Store(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (a *showService) Update(c context.Context, id int64, m *models.Show) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	//Validar que no exista

	err := a.showRepo.Update(ctx, id, m)
	if err != nil {
		return err
	}

	return nil
}

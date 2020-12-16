package service

import (
	"context"
	"week4/internal/model"
	"week4/internal/repository"
)

type Service struct {
	Repo repository.Repository
}

var Srv *Service

func init() {
	Srv, _, _ = New(repository.Repo)
}

func New(repo repository.Repository) (s *Service, cf func(), err error) {
	s = &Service{
		Repo: repo,
	}
	cf = s.Close
	return
}

func (this *Service) GetArticle(ctx context.Context, id int) (model.Article, error) {
	return this.Repo.GetArticle(ctx, id)
}

func (this *Service) Close() {

}

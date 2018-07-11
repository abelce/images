package application

import (
	"admin/domain/model"
	"admin/application/command"
	"admin/application/commandHandler"
)

type service struct {
	Repository    model.Repository
	QueryService  model.QueryService
}

func (s *service) CreateArticle(c command.CreateArticle) (*model.Article) {
	handler := commandHandler.CreateArticle {
		ArticleRepository:    s.Repository,
		QueryService:         s.QueryService,
	}

	article, err := handler.Handle(c)
	if err != nil {
		return nil
	}
	return article
}
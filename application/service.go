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

func (s *service) CreateArticle(c command.CreateArticle) (*model.Article, error) {
	handler := commandHandler.CreateArticle {
		ArticleRepository:    s.Repository,
		QueryService:         s.QueryService,
	}
	return handler.Handle(c)
}

func (s *service) UpdateArticle(c command.UpdateArticle) (*model.Article, error) {
	handler := commandHandler.UpdateArticle{
		ArticleRepository: s.Repository,
		QueryService: s.QueryService,
	}
	return handler.Handle(c)
}

func (s *service) DeleteArticle(c command.DeleteArticle) error {
	handler := commandHandler.DeleteArticle{
		ArticleRepository: s.Repository,
		QueryService: s.QueryService,
	}
	return handler.Handle(c)
}
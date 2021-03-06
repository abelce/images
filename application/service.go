package application

import (
	"images/domain/model"
	"images/application/command"
	"images/application/commandHandler"
)

type service struct {
	Repository    model.Repository
	QueryService  model.QueryService
}

func (s *service) CreateImage(c command.CreateImage) (*model.Image, error) {
	handler := commandHandler.CreateImage {
		ImageRepository:    s.Repository,
		QueryService:         s.QueryService,
	}
	return handler.Handle(c)
}

// func (s *service) UpdateArticle(c command.UpdateArticle) (*model.Article, error) {
// 	handler := commandHandler.UpdateArticle{
// 		ArticleRepository: s.Repository,
// 		QueryService: s.QueryService,
// 	}
// 	return handler.Handle(c)
// }

// func (s *service) DeleteArticle(c command.DeleteArticle) error {
// 	handler := commandHandler.DeleteArticle{
// 		ArticleRepository: s.Repository,
// 		QueryService: s.QueryService,
// 	}
// 	return handler.Handle(c)
// }
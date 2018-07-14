package commandHandler

import (
	"errors"

	"admin/application/command"
	"admin/domain/model"
)

type UpdateArticle struct {
	ArticleRepository model.Repository
	QueryService      model.QueryService
}

func (h UpdateArticle)Handle(c command.UpdateArticle) (*model.Article, error) {
	article, err := h.QueryService.FindByID(c.ID)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("no article found")
	}
	article.Update(c.Title, c.Markdowncontent, c.Private, c.Tags, c.Status, c.Categories, c.Type, c.Description)
	_, err = h.ArticleRepository.UpdateByID(article.ID, article)
	if err != nil {
		return nil, err
	}
	return article, nil
}
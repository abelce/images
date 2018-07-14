package commandHandler 

import (
	"errors"
	
	"admin/domain/model"
	"admin/application/command"
)
type DeleteArticle struct {
	ArticleRepository    model.Repository
	QueryService         model.QueryService
}

func (h DeleteArticle) Handle(c command.DeleteArticle) error {
	article, err := h.QueryService.FindByID(c.ID)
	if err != nil {
		return err
	}
	if article == nil {
		return errors.New("no article found")
	}
	article.Delete()
	_, err = h.ArticleRepository.UpdateByID(article.ID, article)
	if err != nil {
		return err
	}
	return nil
}
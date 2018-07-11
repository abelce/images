package commandHandler

import (
	"admin/application/command"
	"admin/domain/model"
)

type CreateArticle struct {
	ArticleRepository    model.Repository
	QueryService         model.QueryService
}

// (title, markdowncontent, private, tags, status, categories, typ, description string)
func (h CreateArticle) Handle(c command.CreateArticle) (*model.Article, error) {
	
	article := model.NewArticle(c.Title, c.Markdowncontent, c.Private, c.Tags, c.Status, c.Categories, c.Type, c.Description)
    return article, nil;
}
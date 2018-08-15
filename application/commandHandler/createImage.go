package commandHandler

import (
	"admin/application/command"
	"admin/domain/model"
)

type CreateImage struct {
	ImageRepository    model.Repository
	QueryService         model.QueryService
}

// (title, markdowncontent, private, tags, status, categories, typ, description string)
func (h CreateImage) Handle(c command.CreateImage) (*model.Article, error) {
	
	image := model.NewArticle(c.Url)
	err := h.ImageRepository.Save(image)
	if err != nil {
		return nil, err
	}
    return image, nil;
}
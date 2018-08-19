package commandHandler

import (
	"image"
	"os"
	"fmt"

	"images/application/command"
	"images/domain/model"
)

type CreateImage struct {
	ImageRepository    model.Repository
	QueryService         model.QueryService
}

// (title, markdowncontent, private, tags, status, categories, typ, description string)
func (h CreateImage) Handle(c command.CreateImage) (*model.Image, error) {
	
	// 设置图片的宽度，高度
	file, _ := os.Open(c.Url)
	img, str, err := image.DecodeConfig(file)
	fmt.Println(str)
	if err != nil {
		return nil, err
	}
	c.Width = img.Width
	c.Height = img.Height

	image := model.NewImage(c.Url, c.Width, c.Height)
	err = h.ImageRepository.Save(image)
	if err != nil {
		return nil, err
	}
    return image, nil;
}
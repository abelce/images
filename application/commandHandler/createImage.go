package commandHandler

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"os"

	"images/application/command"
	"images/domain/model"
)

type CreateImage struct {
	ImageRepository model.Repository
	QueryService    model.QueryService
}

// (title, markdowncontent, private, tags, status, categories, typ, description string)
func (h CreateImage) Handle(c command.CreateImage) (*model.Image, error) {

	// 设置图片的宽度，高度
	u, err := url.Parse(c.Url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	path := "/data/upload_files" + u.Path
	file, _ := os.Open(path)
	defer file.Close()
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.Width = img.Width
	c.Height = img.Height
	image := model.NewImage(c.Url, c.Width, c.Height)
	err = h.ImageRepository.Save(image)
	if err != nil {
		return nil, err
	}
	return image, nil
}

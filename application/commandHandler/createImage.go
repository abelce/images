package commandHandler

import (
	"fmt"
	"image"
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

	fmt.Println("handler")
	// 设置图片的宽度，高度
	u, err := url.Parse(c.Url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.Path)
	//	path := "/data/upload_files"
	file, _ := os.Open(c.Url)
	fmt.Println(c.Url)
	img, str, err := image.DecodeConfig(file)
	fmt.Println(str)
	if err != nil {
		fmt.Println("err")
		return nil, err
	}
	c.Width = img.Width
	c.Height = img.Height
	fmt.Println(c.Width)
	image := model.NewImage(c.Url, c.Width, c.Height)
	err = h.ImageRepository.Save(image)
	if err != nil {
		return nil, err
	}
	return image, nil
}

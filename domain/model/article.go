package model

import (
	"time"
)

type Article struct {
	ID                    string `json:"id"`
	Title                 string `json:"title"`
	Markdowncontent       string `json:"markdowncontent"`
	Private               string `json:"private"`
	Tags                  string `json:"tags"`
	Status                string `json:"status"`
	Categories            string `json:"categories"`
	Type                  string `json:"type"`   //original
	Description           string `json:"description"`
	CreateTime            int64 `json:"createTime"`		//创建时间
	LastUpdateTime        int64 `json:"lastUpdateTime"`		//最近的更新时间
	Deleted               bool  `json:"deleted"`             
}

func NewArticle(title, markdowncontent, private, tags, status, categories, typ, description string) *Article {
	createTime := time.Now().Unix()
	lastUpdateTime := createTime
	return &Article {
		Title:                 title,
		Markdowncontent:       markdowncontent,
		Private:               private,
		Tags:                  tags,
		Status:                status,
		Categories:            categories,
		Type:                  typ,
		Description:           description,
		Deleted:               false,
		CreateTime:            createTime,
		LastUpdateTime:        lastUpdateTime,
	}
}


func (c *Article) Delete() {
  c.Deleted = true
}


func (c *Article) Modify() {

}
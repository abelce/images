package model

import (
	"time"
)

type Article struct {
	ID                    string `json:"id" jsonapi:"primary,article"`
	Title                 string `json:"title" jsonapi:"attr,title"`
	Markdowncontent       string `json:"markdowncontent" jsonapi:"attr,markdowncontent"`
	Private               string `json:"private" jsonpi:"attr,private"`
	Tags                  string `json:"tags" jsonapi:"attr,tags"`
	Status                string `json:"status" jsonapi:"attr,status"`
	Categories            string `json:"categories" jsonapi:"attr,categories"`
	Type                  string `json:"type" jsonapi:"attr,type"`   //original
	Description           string `json:"description" jsonapi:"attr,description"`
	CreateTime            int64 `json:"createTime,omitempty" jsonapi:"attr,createTime,omitempty"`		//创建时间
	LastUpdateTime        int64 `json:"lastUpdateTime,omitempty" jsonapi:"attr,lastUpdateTime,omitempty"`		//最近的更新时间
	Deleted               bool  `json:"deleted" jsonapi:"attr,deleted"`             
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

func (c *Article) Update (title, markdowncontent, private, tags, status, categories, typ, description string) {
	c.LastUpdateTime = time.Now().Unix()
	c.Title = title
	c.Markdowncontent = markdowncontent
	c.Private = private
	c.Tags = tags
	c.Status = status
	c.Categories = categories
	c.Type = typ
	c.Description = description
}

func (c *Article) Delete() {
    c.LastUpdateTime = time.Now().Unix()
    c.Deleted = true
}


func (c *Article) Modify() {

}
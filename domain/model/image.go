package model

import (
	"time"
)

type Image struct {
	ID                    string `json:"id" jsonapi:"primary,image"`
	Url                   string `json:"url" jsonapi:"attr,url"`  
	Deleted               bool `json:"deleted" jsonapi:"attr,deleted"`  
	CreateTime            int64 `json:"createTime" jsonapi:"attr,createTime"`  
	LastUpdateTime        int64 `json:"lastUpdateTime" jsonapi:"attr,lastUpdateTime"`          
}

func NewArticle(url string) *Image {
	createTime := time.Now().Unix()
	lastUpdateTime := createTime
	return &Image {
		Url:                   url,
		Deleted:               false,
		CreateTime:            createTime,
		LastUpdateTime:        lastUpdateTime,
	}
}

func (c *Image) Delete() {
    c.LastUpdateTime = time.Now().Unix()
    c.Deleted = true
}


func (c *Image) Modify() {

}
package model

import (
	"time"
)

type Image struct {
	ID                    string `json:"id" jsonapi:"primary,image"`
	Url                   string `json:"url" jsonapi:"attr,url"`  
	Width                 int `json:"width" jsonapi:"attr,width"`
	Height                int `json:"height" jsonapi:"attr,height"`
	Deleted               bool `json:"deleted" jsonapi:"attr,deleted"`  
	CreateTime            int64 `json:"createTime" jsonapi:"attr,createTime"`  
	LastUpdateTime        int64 `json:"lastUpdateTime" jsonapi:"attr,lastUpdateTime"`          
}

func NewImage(url string, width int, height int) *Image {
	createTime := time.Now().Unix()
	lastUpdateTime := createTime
	return &Image {
		Url:                   url,
		Width:                 width,
		Height:                height,
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
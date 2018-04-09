package domain

import ()

type User struct {
	ID 			string `json:"id"`
	Email		string `json:"email"`
	Password	string `json:"password"`
	FirstName 	string `json:"firstName"`
	LastName  	string `json:"lastName"`
	Phone		string `json:"phone"`
	Sex			string `json:"sex"`
	IsAdmin		bool `json:"isAdmin"`
	CreateTime  int64	`json:"createTime"`		//创建时间
	LastUpdateTime int64  `json:"lastUpdateTime"`		//最近的更新时间
	AccessTime  int64	`json:"accessTime"`		//最近登陆时间
}
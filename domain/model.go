package domain

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
	LogoImage   string   `json:"logoImage"`     //logo url
}

type Article struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Markdowncontent string `json:"markdowncontent"`
	Private    string `json:"private"`
	Tags       string `json:"tags"`
	Status     string `json:"status"`
	Categories string  `json:"categories"`
	Type       string  `json:"type"`   //original
	Description string  `json:"description"`
	CreateTime  int64	`json:"createTime"`		//创建时间
	LastUpdateTime int64  `json:"lastUpdateTime"`		//最近的更新时间
}
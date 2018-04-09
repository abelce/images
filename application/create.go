package application

type Create struct {
	ID 			string `json:"id"`
	Email		string `json:"email"`
	Password	string `json:"password"`
	FirstName 	string `json:"firstName"`
	LastName  	string `json:"lastName"`
	Phone		string `json:"phone"`
	Sex			string `json:"sex"`
	IsAdmin		string `json:"isAdmin"`
}


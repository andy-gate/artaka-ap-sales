package models

type UserLogin struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type LoginResponse struct {
	Status			bool			`json:"status"`
	Message			string			`json:"message"`
	User			UserResponse	`json:"user"`
	Token			string			`json:"token"`
}

type UserResponse struct {
	Fullname        uint			`json:"fullname"`
	Store			[]StoreDetail	`json:"store"`
}

type StoreDetail struct {
	Store_id        string		`json:"store_id"`
	Airport_code	string		`json:"airport_code"`
	Store_name      string		`json:"store_name"`
	Store_reference	string		`json:"store_reference"`
}
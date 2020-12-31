package domain

type User struct {
	Id 			uint64 `json:"id" xml:"id"`
	FirstName 	string `json:"first_name" xml:"first_name"`
	LastName	string `json:"last_name" xml:"last_name"`
	Email 		string `json:"email" xml:"email"`
}


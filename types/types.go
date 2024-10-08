package types

import "time"

type UserStore interface{
	GetUserByEmail(email string) (*User,error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type mockUserStore struct{

}

func GetUserByEmail(email string) (*User,error){
	return nil,nil
}

type User struct{
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string	`json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string	`json:"password"`
}
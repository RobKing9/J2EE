package models

type User struct {
	Username string `json:"username"`
	Psw      string `json:"psw"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Sex      string `json:"sex"`
	Borndate string `json:"borndate"`
}

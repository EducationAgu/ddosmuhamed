package model

type User struct {
	Id       int64
	Name     string
	Username string
	Email    string
	Password string `json:"password,omitempty"`

	Salt string
}

package models

// User is an user on go-bancho.
type User struct {
	ID          int
	Username    string `sql:"size:20"`
	Permissions uint32
	Email       string
	Password    string
}
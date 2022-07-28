package models

import (
	"encoding/json"
)

var Userdata string = `[{"name": "Raj","age": "31","address": {"street": "My street","city": "US city","state": "Us State"}},{"name": "Nimmy","age": "31","address": {"street": "My street","city": "US city","state": "Us State"}}]`

var Users []User

type Userserviceinterface interface {
	GetUsers() *[]User
	Createuser(user *User) *[]User
}

func Newmodel() Userserviceinterface {
	Loadjson()
	return &User{}
}

func Loadjson() *[]User {
	json.Unmarshal([]byte(Userdata), &Users)
	return &Users
}

func (Myuser *User) GetUsers() *[]User {
	json.Unmarshal([]byte(Userdata), &Users)
	return &Users
}

func (userlist *User) Createuser(user *User) *[]User {
	Users = append(Users, *user)
	return &Users
}

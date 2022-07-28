package models

type UserInterface interface {
	Createuser(User) User
	GetAllusers() []User
}

type UserService struct {
	Users []User
}

func New() UserInterface {
	return &UserService{}
}

func getData() UserInterface {
	return &UserService{}
}

func (us *UserService) Createuser(user User) User {
	us.Users = append(us.Users, user)
	return user
}

func (us *UserService) GetAllusers() []User {
	return us.Users
}

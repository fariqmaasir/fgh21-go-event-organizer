package models

import (
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`   //binding:"required"`
	Email    string `json:"email" form:"email"` //binding:"required,email"`
	Password string `json:"-" form:"password"`  //validate:"required"`
}

var data []User = []User{
	{
		Id:       1,
		Name:     "fazz",
		Email:    "fazz@mail.com",
		Password: "12345678",
	},
	{
		Id:       2,
		Name:     "kiko",
		Email:    "kiko@mail.com",
		Password: "kikooooo",
	},
	{
		Id:       3,
		Name:     "lana",
		Email:    "lana@mail.com",
		Password: "5678",
	},
}

func FindAllUsers() []User {
	return data
}

func FindOneUser(id int) any {
	fmt.Println(data)
	fmt.Println(id)
	user := []User{}
	for _, acc := range data {
		if acc.Id == id {
			user = []User{
				{
					Id:       acc.Id,
					Name:     acc.Name,
					Email:    acc.Email,
					Password: acc.Password,
				},
			}
		}
	}
	return user
}

func CreateOneUser(user User) []User {
	data = append(data, User{
		Id:       len(data) + 1,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	// fmt.Println(len(data))
	return data
}

func DeleteOneUser(id int) []User {
	data := data
	for i, x := range data {
		if x.Id == id {
			data = append(data[:i], data[i+1:]...)
		}
	}
	return data
}
func UpdateOneUser(id int, user User) []User {
	data := data
	for i, acc := range data {
		if acc.Id == id {
			data[i].Name = user.Name
			data[i].Email = user.Email
			data[i].Password = user.Password
		}
	}
	return data
}

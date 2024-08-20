package models

import (
	"context"
	"fmt"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int     `json:"id"`
	Email    string  `json:"email" form:"email"`       //binding:"required,email"`
	Password string  `json:"-" form:"password"`        //validate:"required"`
	Username *string `json:"username" form:"username"` //binding:"required"`
}

var data []User = []User{
	{
		Id:       1,
		Email:    "fazz@mail.com",
		Password: "12345678",
		// Username: "fazz",
	},
	{
		Id:       2,
		Email:    "kiko@mail.com",
		Password: "kikooooo",
		// Username: "kiko",
	},
	{
		Id:       3,
		Email:    "lana@mail.com",
		Password: "5678",
		// Username: "lana",
	},
}

func PagesInfos(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	var table int
	err := db.QueryRow(
		context.Background(),
		`select count (id) as table where "title" ilike '%' || ($1) || '%'  from "users"`, search,
	).Scan(&table)
	fmt.Println(table)
	if err != nil {
		fmt.Print(err)
	}
	return table
}

func FindAllUsers(search string, page int, limit int) ([]User, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	totalData := PagesInfo(search)
	var offset int = (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`
		select * from "users" 
		where "email" 
		ilike '%' || ($1) || '%'
		limit ($2)
		offset ($3)
		`,
		search,
		limit,
		offset,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])
	if err != nil {
		fmt.Println(err)
	}
	return users, totalData
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
					Email:    acc.Email,
					Password: acc.Password,
					Username: acc.Username,
				},
			}
		}
	}
	return user
}

func CreateOneUser(user User) User {
	db := lib.DB()
	defer db.Close(context.Background())

	//encrypt
	user.Password = lib.Encrypt(user.Password)
	//encrypt

	fmt.Println(user)
	sql := `insert into "users" ("email", "password", "username") values ($1, $2, $3) returning "id","email","password","username"`
	row := db.QueryRow(context.Background(), sql, user.Email, user.Password, user.Username)

	var results User
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.Username,
	)

	fmt.Println(results)
	fmt.Println(row)
	return results
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
			data[i].Email = user.Email
			data[i].Password = user.Password
			data[i].Username = user.Username
		}
	}
	return data
}
func FindOneUserByEmail(email string) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(
		context.Background(),
		`select * from "users"`,
	)

	fmt.Println(rows)
	fmt.Println(err)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])
	fmt.Println(users)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	user := User{}
	for _, acc := range users {
		if acc.Email == email {
			user = acc
		}
	}
	return user
}

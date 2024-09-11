package models

import (
	"context"
	"fmt"
	"time"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id            int        `json:"id" db:"id"`
	Picture       *string    `json:"picture" form:"picture" db:"picture"`
	FullName      *string    `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate     *time.Time `json:"birthDate" form:"birthDate" db:"birth_date"`
	Gender        *int       `json:"gender" form:"gender" db:"gender"`
	PhoneNumber   *string    `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession    *string    `json:"profession" form:"profession" db:"profession"`
	NationalityId *int       `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId        int        `json:"userId" form:"userId" db:"user_id"`
}

type Regist struct {
	Id            int        `json:"id"`
	Email         string     `json:"email" form:"email" db:"email"`
	Password      string     `json:"-" form:"password" db:"password"`
	Username      *string    `json:"username" form:"username"`
	Picture       *string    `json:"picture" form:"picture" db:"picture"`
	FullName      *string    `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate     *time.Time `json:"birthDate" form:"birthDate" db:"birth_date"`
	Gender        *int       `json:"gender" form:"gender" db:"gender"`
	PhoneNumber   *string    `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession    *string    `json:"profession" form:"profession" db:"profession"`
	NationalityId *int       `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId        int        `json:"userId" form:"userId" db:"user_id"`
}

func CreateProfile(regist Regist) (*Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	regist.Password = lib.Encrypt(regist.Password)
	var userId int
	err := db.QueryRow(
		context.Background(),
		`INSERT INTO "users" ("email", "password", "username") VALUES ($1, $2, $3) RETURNING "id"`,
		regist.Email, regist.Password, regist.Username,
	).Scan(&userId)

	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into users table")
	}

	var profile Profile
	picture := "http://localhost:8888/images/profile.jpg"
	err = db.QueryRow(
		context.Background(),
		`INSERT INTO "profile" ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id", "user_id") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id`,
		picture, regist.FullName, regist.BirthDate, regist.Gender,
		regist.PhoneNumber, regist.Profession, regist.NationalityId, userId,
	).Scan(
		&profile.Id,
		&profile.Picture,
		&profile.FullName,
		&profile.BirthDate,
		&profile.Gender,
		&profile.PhoneNumber,
		&profile.Profession,
		&profile.NationalityId,
		&profile.UserId,
	)

	fmt.Println(err)
	fmt.Println("---")
	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table")
	}

	return &profile, nil
}

func UpdateProfileImage(data Profile, id int) (Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET picture = $1 WHERE user_id=$2 returning *`

	row, err := db.Query(context.Background(), sql, data.Picture, id)
	if err != nil {
		return Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Profile])
	if err != nil {
		return Profile{}, nil
	}

	return profile, nil
}

func FindOneProfile(id int) (*Regist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var profile Regist
	err := db.QueryRow(context.Background(), `select profile.id, users.email, users.username, profile.picture, profile.full_name, profile.birth_date, profile.gender, profile.phone_number, profile.profession, profile.nationality_id from profile inner join users on profile.user_id = users.id where user_id = $1`, id).Scan(
		&profile.Id,
		&profile.Email,
		&profile.Username,
		&profile.Picture,
		&profile.FullName,
		&profile.BirthDate,
		&profile.Gender,
		&profile.PhoneNumber,
		&profile.Profession,
		&profile.NationalityId,
	)

	fmt.Println(err)
	fmt.Println("---")
	if err != nil {
		return nil, fmt.Errorf("acc not found")
	}
	return &profile, err
}

func EditProfile(regist Regist, id int) (*Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	regist.Password = lib.Encrypt(regist.Password)
	var userId int
	err := db.QueryRow(
		context.Background(),
		`update "users" set ("email", username) = ($1, $2) where id = $3 RETURNING "id"`,
		regist.Email, regist.Username, id,
	).Scan(&userId)

	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into users table")
	}

	var profile Profile
	err = db.QueryRow(
		context.Background(),
		`update "profile" set ("full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id") 
		= ($1, $2, $3, $4, $5, $6) where user_id = $7 RETURNING id, picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id`,
		regist.FullName, regist.BirthDate, regist.Gender,
		regist.PhoneNumber, regist.Profession, regist.NationalityId, id,
	).Scan(
		&profile.Id,
		&profile.Picture,
		&profile.FullName,
		&profile.BirthDate,
		&profile.Gender,
		&profile.PhoneNumber,
		&profile.Profession,
		&profile.NationalityId,
		&profile.UserId,
	)

	fmt.Println(err)
	fmt.Println("---")
	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table")
	}

	return &profile, nil
}

type ChangePass struct {
	Id          int
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}

func ChangePassword(Pass ChangePass, id int) (int, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	Pass.NewPassword = lib.Encrypt(Pass.NewPassword)

	var password string
	err1 := db.QueryRow(context.Background(), `select "password" from users where id = $1`, id).Scan(&password)
	// fmt.Print(err1)
	if err1 != nil {
		return 0, fmt.Errorf("password doesnt exist")
	}
	fmt.Println(password)
	fmt.Println(Pass.OldPassword)
	isVerified := lib.Verify(Pass.OldPassword, password)
	fmt.Println(isVerified)
	fmt.Println("-")
	if !isVerified {
		return 0, fmt.Errorf("wrong password")
	}

	dataSql := `update "users" set "password" = $1 where id = $2`
	_, err := db.Exec(context.Background(), dataSql, Pass.NewPassword, id)
	// fmt.Println(err)
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert")
	}

	return id, nil
}

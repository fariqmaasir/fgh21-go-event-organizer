package models

import (
	"context"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Nationality struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ListAllNationalities() ([]Nationality, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), `select * from nationalities`)

	nationality, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Nationality])
	if err != nil {
		return nil, err
	}
	return nationality, nil
}

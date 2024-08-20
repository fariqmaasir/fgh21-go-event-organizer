package models

import (
	"context"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Partners struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

func ListAllPartners() ([]Partners, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), `select * from partners`)

	partners, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Partners])
	if err != nil {
		return nil, err
	}
	return partners, nil
}

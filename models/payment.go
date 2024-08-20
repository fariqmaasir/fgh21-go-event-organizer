package models

import (
	"context"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type PaymentMethod struct {
	Id   int    `json:"id"`
	Name string `json:"paymentMethod"`
}

func ListAllPaymentMethod() ([]PaymentMethod, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), `select * from payment_method`)

	payment, err := pgx.CollectRows(rows, pgx.RowToStructByPos[PaymentMethod])
	if err != nil {
		return nil, err
	}
	return payment, nil
}

package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(),
		"postgresql://postgres:32788a444163454cac61c1d26a0cdfb1@localhost:5432/event_organizer?sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}

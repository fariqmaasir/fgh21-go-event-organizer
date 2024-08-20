package models

import (
	"context"
	"fmt"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type EventCategory struct {
	Id         int
	EventId    int
	CategoryId int
}

func CreateOneEventCategory(category EventCategory) (*EventCategory, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	// fmt.Println(id)
	var eventCategory EventCategory
	err := db.QueryRow(
		context.Background(),
		`insert into "eventCategory" ("event_category", "category_id") values ($1, $2) returning "id", "event_category", "category_id"`,
		category.EventId, category.CategoryId,
	).Scan(
		&eventCategory.Id, &eventCategory.EventId, &eventCategory.CategoryId,
	)

	fmt.Println(eventCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert")
	}

	return &eventCategory, nil
}

func PageInfo(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	var table int
	err := db.QueryRow(
		context.Background(),
		`select count (id) as table where "category_id" ilike '%' || ($1) || '%'  from "eventCategory"`, search,
	).Scan(&table)
	fmt.Println(table)
	if err != nil {
		fmt.Print(err)
	}
	return table
}

func FindAllEventCategory(search string, page int, limit int) ([]EventCategory, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	totalData := PageInfo(search)
	var offset int = (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`
		select * from "eventCategory" 
		where "category_id" 
		ilike '%' || ($1) || '%'
		limit ($2)
		offset ($3)
		`,
		search,
		limit,
		offset,
	)

	eventCategory, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventCategory])
	if err != nil {
		fmt.Println(err)
	}
	return eventCategory, totalData
}

func EditOneEventCategory(eventCategory EventCategory, id int, createId int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "eventCategory" set ("event_category", "category_id") = ($1, $2) where id=$7`

	fmt.Println(eventCategory)
	_, err := db.Exec(context.Background(), dataSql, eventCategory.EventId, eventCategory.CategoryId, createId, id)
	fmt.Println(err)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func DeleteOneEventCategory(id int) (*EventCategory, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var eventCategory EventCategory
	err := db.QueryRow(context.Background(), "delete from eventCategory where id = ($1) returning id, event_id, category_id, date, descriptions, location_id, created_by", id).
		Scan(
			&eventCategory.Id,
			&eventCategory.EventId,
			&eventCategory.CategoryId,
		)

	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table")
	}
	return &eventCategory, nil
}

func FindOneEventCategory(id int) []EventCategory {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		"select * from eventCategory where id = ($1)",
		id,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventCategory])

	fmt.Println(rows.Values())
	fmt.Println("")
	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	}
	return users
}

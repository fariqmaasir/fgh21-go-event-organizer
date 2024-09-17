package models

import (
	"context"
	"fmt"
	"time"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id           *int      `json:"id" db:"id"`
	Image        *string   `json:"image" form:"image" db:"image"`
	Title        *string   `json:"title" form:"title" db:"title"`
	Date         time.Time `json:"date" form:"date" db:"date"`
	Descriptions *string   `json:"descriptions" form:"descriptions" db:"descriptions"`
	LocationId   *int      `json:"locationId" form:"locationId" db:"location_id"`
	CreatedBy    int       `json:"createdBy" db:"created_by"`
}

type EventSection struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Price    int    `json:"price" db:"price"`
	Quantity int    `json:"quantity" db:"quantity"`
	EventId  int    `json:"eventId" db:"event_id"`
}

func CreateOneEvent(event Events, id int) (*Events, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var events Events
	err := db.QueryRow(
		context.Background(),
		`insert into "events" ("image", "title", "date", "descriptions", "location_id", "created_by") values ($1, $2, $3, $4, $5, $6) returning "id", "image", "title", "date", "descriptions", "location_id", "created_by"`,
		event.Image, event.Title, event.Date, event.Descriptions, event.LocationId, id,
	).Scan(
		&events.Id, &events.Image, &events.Title, &events.Date, &events.Descriptions, &events.LocationId, &events.CreatedBy,
	)

	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert")
	}

	return &events, nil
}

func PagesInfo(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	var table int
	err := db.QueryRow(
		context.Background(),
		`select count (id) as table where "title" ilike '%' || ($1) || '%'  from "events"`, search,
	).Scan(&table)
	fmt.Println(table)
	if err != nil {
		fmt.Print(err)
	}
	return table
}

func FindAllEvents(search string, page int, limit int) ([]Events, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	totalData := PagesInfo(search)
	var offset int = (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`
		SELECT * from "events" 
		WHERE "title" 
		ilike '%' || ($1) || '%'
		limit ($2)
		offset ($3)
		`,
		search,
		limit,
		offset,
	)

	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(totalData)
	return events, totalData
}

func EditOneEvent(event Events, id int, createId int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "events" set ("image", "title", "date", "descriptions", "location_id", "created_by") = ($1, $2, $3, $4, $5, $6) where id=$7`

	fmt.Println(event)
	_, err := db.Exec(context.Background(), dataSql, event.Image, event.Title, event.Date, event.Descriptions, event.LocationId, createId, id)
	fmt.Println(err)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func DeleteOneEvent(id int) (*Events, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var event Events
	err := db.QueryRow(context.Background(), "delete from events where id = ($1) returning id, image, title, date, descriptions, location_id, created_by", id).
		Scan(
			&event.Id,
			&event.Image,
			&event.Title,
			&event.Date,
			&event.Descriptions,
			&event.LocationId,
			&event.CreatedBy,
		)

	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table")
	}
	return &event, nil
}

func FindOneEvent(id int) Events {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		"select * from events where id = ($1)",
		id,
	)

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Events])

	if err != nil {
		fmt.Println(err)
	}
	return users
}

func FindOneEventByUserId(id int) []Events {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		"select * from events where created_by = ($1)",
		id,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])

	if err != nil {
		fmt.Println(err)
	}
	return users
}

func FindSectionsByEventId(id int) ([]EventSection, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), `select * from "event_sections" where event_id = $1`, id)

	section, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventSection])

	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return section, nil
}

type Wishlist struct {
	Id      int `json:"id"`
	EventId int `json:"eventId" db:"event_id" form:"eventId"`
	UserId  int `json:"userId" db:"user_id"`
}

func CreateOneWishlist(eventId int, userId int) (*Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var wish Wishlist
	err := db.QueryRow(
		context.Background(),
		`insert into "wishlist" ("event_id", "user_id") values ($1, $2) returning "id", "event_id", "user_id"`,
		eventId, userId,
	).Scan(
		&wish.Id, &wish.EventId, &wish.UserId,
	)

	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert")
	}

	return &wish, nil
}
func DeleteWishlistById(id int) (*Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var wish Wishlist
	err := db.QueryRow(context.Background(), `delete from wishlist where id = ($1) returning "id", "event_id", "user_id"`, id).Scan(&wish.Id, &wish.EventId, &wish.UserId)

	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("error when delete")
	}
	return &wish, nil
}

type EventWishlist struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	// Location string    `json:"location"`
	UserId int `json:"userId"`
}

func GetWishlistByUserId(id int) ([]EventWishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select wishlist.id, events.title, events.date, wishlist.user_id
		from wishlist
		inner join users on wishlist.user_id = users.id
		inner join events on wishlist.event_id = events.id
		where users.id = $1 LIMIT 100`,
		id,
	)

	wish, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventWishlist])

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	fmt.Println(wish)
	return wish, nil
}

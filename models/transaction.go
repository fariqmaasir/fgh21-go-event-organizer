package models

import (
	"context"
	"fmt"
	"time"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Transactions struct {
	Id             int   `json:"id"`
	EventId        int   `json:"eventId" form:"eventId" db:"event_id"`
	PaymentId      int   `json:"paymentId" form:"paymentId" db:"payment_method_id"`
	UserId         int   `json:"userId"  db:"user_id"`
	SectionId      []int `json:"sectionId,omitempty" form:"sectionId" db:"section_id"`
	TicketQuantity []int `json:"ticketQuantity,omitempty" form:"ticketQuantity" db:"ticket_qty"`
}

// func CreateTransaction(form Transactions, id int) (*Transactions, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())
// 	fmt.Println(id)
// 	var transaction Transactions
// 	var transactionId int
// 	err := db.QueryRow(
// 		context.Background(),
// 		`insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`,
// 		form.EventId, form.PaymentId, id,
// 	).Scan(
// 		&transactionId, &transaction.EventId, &transaction.PaymentId, &transaction.UserId,
// 	)
// 	fmt.Println(err)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute insert")
// 	}
// 	err1 := db.QueryRow(
// 		context.Background(),
// 		`insert into "transaction_details" ("transaction_id", "section_id", "ticket_qty") values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`,
// 		transactionId, form.SectionId, form.TicketQuantity,
// 	).Scan(
// 		&transaction.Id, &transaction.PaymentId, &transaction.SectionId, &transaction.TicketQuantity,
// 	)
// 	fmt.Println(err1)
// 	if err1 != nil {
// 		return nil, fmt.Errorf("failed to execute insert")
// 	}
// 	return &transaction, nil
// }

func CreateNewTransactions(data Transactions) Transactions {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	row := db.QueryRow(context.Background(), sql, data.EventId, data.PaymentId, data.UserId)

	var results Transactions
	row.Scan(
		&results.Id,
		&results.EventId,
		&results.PaymentId,
		&results.UserId,
	)
	fmt.Println(results)
	return results
}

type TransactionDetail struct {
	Id             int `json:"id"`
	TransactionId  int `json:"transactionId" form:"transactionId" db:"transaction_id"`
	SectionId      int `json:"sectionId" form:"sectionId" db:"section_id"`
	TicketQuantity int `json:"ticketQuantity" form:"ticketQuantity" db:"ticket_qty"`
}

func CreateTransactionDetail(data TransactionDetail) TransactionDetail {
	db := lib.DB()
	defer db.Close(context.Background())

	inputSQL := `insert into "transaction_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
	row := db.QueryRow(context.Background(), inputSQL, data.TransactionId, data.SectionId, data.TicketQuantity)

	var detail TransactionDetail

	row.Scan(
		&detail.Id,
		&detail.TransactionId,
		&detail.SectionId,
		&detail.TicketQuantity,
	)
	fmt.Println(row)
	return detail
}

type Result struct {
	Id             int       `json:"id"`
	FullName       string    `json:"fullName" form:"fullName" db:"full_name"`
	EventTitle     string    `json:"eventTitle" form:"eventTitle" db:"title"`
	LocationId     *int      `json:"location_id" form:"location_id" db:"location"`
	Date           time.Time `json:"date" form:"date" db:"date"`
	PaymentId      string    `json:"PaymentId" form:"PaymentId" db:"payment_method_id"`
	SectionName    []string  `json:"sectionName" form:"sectionName" db:"name"`
	TicketQuantity []int     `json:"TicketQuantity" form:"TicketQuantity" db:"ticket_qty"`
}

func AddDetailsTransaction() ([]Result, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql :=
		`select t.id, p.full_name, e.title as "event_title", e.location_id, e.date, pm.name as "payment_method",
        array_agg(es.name) as "section_name", array_agg(td.ticket_qty) as "ticket_qty"
        from "transactions" "t"
        join "users" "u" on u.id = t.user_id
        join "profile" "p" on p.user_id = u.id
        join "events" "e" on e.id = t.event_id
        join "payment_method" "pm" on pm.id = t.payment_method_id
        join "transaction_details" "td" on td.transaction_id = t.id
        join "event_sections" "es" on es.id = td.section_id
        group by t.id, p.full_name, e.title, e.location_id, e.date, pm.name;`

	send, _ := db.Query(
		context.Background(),
		sql,
	)

	row, err := pgx.CollectRows(send, pgx.RowToStructByPos[Result])
	if err != nil {
		fmt.Println(err)
	}
	return row, err
}

func ListAllTransactionById(id int) ([]Result, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql :=
		`select t.id, p.full_name, e.title as "event_title", e.location_id, e.date, pm.name as "payment_method",
        array_agg(es.name) as "section_name", array_agg(td.ticket_qty) as "ticket_qty"
        from "transactions" "t"
        join "users" "u" on u.id = t.user_id
        join "profile" "p" on p.user_id = u.id
        join "events" "e" on e.id = t.event_id
        join "payment_method" "pm" on pm.id = t.payment_method_id
        join "transaction_details" "td" on td.transaction_id = t.id
        join "event_sections" "es" on es.id = td.section_id
		where t.user_id = $1
        group by t.id, p.full_name, e.title, e.location_id, e.date, pm.name`

	send, _ := db.Query(
		context.Background(),
		sql,
		id,
	)

	row, err := pgx.CollectRows(send, pgx.RowToStructByPos[Result])
	if err != nil {
		fmt.Println(err)
	}
	return row, err
}

// func ListOneTransaction(id int) (*Result, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())
// 	var results Result
// 	err := db.QueryRow(
// 		context.Background(),
// 		`
// 		SELECT transaction_details.id, events.image, events.title, events.date, events.descriptions, events.location_id,
// 		events.created_by, payment_method.name, users.email, event_sections.name, event_sections.price, event_sections.quantity FROM transaction_details
// 		inner join transactions on transaction_details.transaction_id = transactions.id
// 		inner join events on transactions.event_id = events.id
// 		inner join payment_method on transactions.payment_method_id = payment_method.id
// 		inner join users on transactions.user_id = users.id
// 		inner join event_sections on transaction_details.section_id = event_sections.id
// 		where transaction_details.id = $1
// 		`,
// 		id,
// 	).Scan(&results.Id, &results.Image, &results.Title, &results.Date, &results.Descriptions, &results.LocationId, &results.CreatedBy, &results.PaymentMethod, &results.Email, &results.Section, &results.Price, &results.Quantity)
// 	fmt.Println(err)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute insert")
// 	}
// 	return &results, nil
// }

type UserTransaction struct {
	Id            int       `json:"id"`
	UserName      *string   `json:"username"`
	Title         string    `json:"title"`
	LocationId    *int      `json:"locationId"`
	Date          time.Time `json:"date"`
	PaymentMethod string    `json:"paymentMethod"`
	Section       string    `json:"section"`
	Quantity      int       `json:"quantity"`
}

// func ListAllTransactionById(id int) ([]UserTransaction, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	// var transaction UserTransaction
// 	rows, _ := db.Query(
// 		context.Background(),
// 		`SELECT transaction_details.id, users.username, events.title, events.location_id, events.date, payment_method.name, event_sections.name, event_sections.quantity FROM transaction_details
// 		inner join transactions on transaction_details.transaction_id = transactions.id
// 		inner join users on transactions.user_id = user_id
// 		inner join events on transactions.event_id =events.id
// 		inner join payment_method on transactions.payment_method_id = payment_method.id
// 		inner join event_sections on transaction_details.section_id = event_sections.id
// 		where transactions.user_id = $1`,
// 		id,
// 	)
// 	transaction, err := pgx.CollectRows(rows, pgx.RowToStructByPos[UserTransaction])
// 	// Scan(&transaction.Id, &transaction.FullName, &transaction.Title, &transaction.LocationId, &transaction.Date, &transaction.PaymentMethod, &transaction.Section, &transaction.Quantity)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transaction, nil
// }

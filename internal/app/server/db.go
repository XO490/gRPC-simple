package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type ordering struct {
	asc  string
	desc string
}

var order = ordering{
	asc:  "ASC",
	desc: "DESC",
}

type product struct {
	Id          string    `pg:"id"`
	Name        string    `pg:"name"`
	CreatedAt   time.Time `pg:"created_at"`
	Price       string    `pg:"price"`
	Description string    `pg:"description"`
}

func conn() *sql.DB {
	db, err := sql.Open("postgres", "host='localhost' port=5432 user=grpcSimple password=grpcSimple dbname=grpcSimple sslmode=disable")
	CheckErr(err, "conn")
	CheckErr(db.Ping(), "conn")
	return db
}

func GetProduct(field string, value string) (*product, error) {
	db := conn()

	defer db.Close()

	query := fmt.Sprintf(`SELECT * FROM "grpcSimple".public."Products" WHERE %s='%s';`, field, value)
	rows, err := db.Query(query)
	CheckErr(err, "GetProduct query")

	p := &product{}
	for rows.Next() {
		row := rows.Scan(
			&p.Id,
			&p.Name,
			&p.CreatedAt,
			&p.Price,
			&p.Description)
		CheckErr(row, "GetProduct rows.Scan")
	}

	return p, nil
}

func GetAllProducts(field string, order string) ([]product, error) {
	db := conn()

	defer db.Close()

	query := fmt.Sprintf(`SELECT * FROM "grpcSimple".public."Products" ORDER BY %s %s;`, field, order)
	rows, err := db.Query(query)
	CheckErr(err, "GetAllProduct query")

	var items []product
	p := &product{}
	for rows.Next() {
		row := rows.Scan(
			&p.Id,
			&p.Name,
			&p.CreatedAt,
			&p.Price,
			&p.Description)
		CheckErr(row, "GetAllProduct rows.Scan")
		items = append(items, *p)
	}

	return items, nil
}

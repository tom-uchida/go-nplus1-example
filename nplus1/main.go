package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	limit, _ := strconv.Atoi(os.Args[1])

	db, _ := sql.Open(
		"postgres",
		"host=localhost port=5432 user=test password=test dbname=testdb sslmode=disable",
	)

	start := time.Now()

	rows, _ := db.Query(
		"SELECT id FROM users LIMIT $1",
		limit,
	)

	for rows.Next() {
		var id int
		rows.Scan(&id)

		postRows, _ := db.Query(
			"SELECT title FROM posts WHERE user_id=$1",
			id,
		)

		for postRows.Next() {
			var title string
			postRows.Scan(&title)
		}
	}

	fmt.Println("elapsed:", time.Since(start))
}

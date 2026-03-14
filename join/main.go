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
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./join <user_count>")
		return
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=test password=test dbname=testdb sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	query := `
	SELECT users.id, users.name, posts.title
	FROM users
	JOIN posts ON users.id = posts.user_id
	WHERE users.id IN (
		SELECT id FROM users LIMIT $1
	)
	`

	rows, err := db.Query(query, limit)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var title string

		rows.Scan(&id, &name, &title)
	}

	fmt.Println("elapsed:", time.Since(start))
}

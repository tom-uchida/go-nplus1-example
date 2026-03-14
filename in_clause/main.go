package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./in_clause <user_count>")
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

	// users取得
	rows, err := db.Query(
		"SELECT id FROM users LIMIT $1",
		limit,
	)
	if err != nil {
		panic(err)
	}

	var ids []int

	for rows.Next() {
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}

	// IN句のプレースホルダ作成
	placeholders := []string{}
	args := []interface{}{}

	for i, id := range ids {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		args = append(args, id)
	}

	query := fmt.Sprintf(
		"SELECT user_id, title FROM posts WHERE user_id IN (%s)",
		strings.Join(placeholders, ","),
	)

	postRows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}

	for postRows.Next() {
		var userID int
		var title string
		postRows.Scan(&userID, &title)
	}

	fmt.Println("elapsed:", time.Since(start))
}

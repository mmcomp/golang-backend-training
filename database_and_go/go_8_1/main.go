package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	go_log "github.com/mmcomp/go-log"
)

func main() {
	db, err := sql.Open("postgres",
		"mehrdad@localhost:3306/mehrdad")
	if err != nil {
		go_log.Log(err)
	}

	// var (
	// 	id           int
	// 	name         string
	// 	email        string
	// 	when_created string
	// )
	rows, err := db.Query("select id, name, email, when_created from accounts")
	if err != nil {
		go_log.Log("Query Error", err)
	}
	defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name, &email, &when_created)
	// 	if err != nil {
	// 		go_log.Log(err)
	// 	}
	// 	go_log.Log(id, name, email, when_created)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	go_log.Log(err)
	// }

	defer db.Close()
}

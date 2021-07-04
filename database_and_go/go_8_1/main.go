package main

import (
	"database/sql"

	"github.com/lib/pq"
	go_log "github.com/mmcomp/go-log"
)

func main() {
	connector, err := pq.NewConnector("user=mehrdad password=123456 dbname=mehrdad")
	if err != nil {
		go_log.Logf("Error Connector : %s", err.Error())
		return
	}
	db := sql.OpenDB(connector)
	rows, err := db.Query("select * from accounts")
	if err != nil {
		go_log.Logf("Error Query : %s", err.Error())
		return
	}
	var id int64
	var name string
	var email string
	var when_created string
	for rows.Next() {
		rows.Scan(&id, &name, &email, &when_created)
		go_log.Logf("id = %d, name = %q, email = %q, when_created = %q", id, name, email, when_created)
	}
	defer db.Close()
}

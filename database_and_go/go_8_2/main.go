package main

import (
	"database/sql"

	"github.com/lib/pq"
	go_log "github.com/mmcomp/go-log"
)

func main() {
	connector, err := pq.NewConnector("user=mehrdad password=123456 dbname=mehrdad sslmode=disable")
	if err != nil {
		go_log.Logf("Error Connector : %s", err.Error())
		return
	}
	db := sql.OpenDB(connector)
	var id int64 = 1
	rows, err := db.Query("select * from accounts where id = $1", id)
	if err != nil {
		go_log.Logf("Error Query : %s", err.Error())
		return
	}
	var selectedId int64
	var name string
	var email string
	var when_created string
	for rows.Next() {
		rows.Scan(&selectedId, &name, &email, &when_created)
		go_log.Logf("id = %d, name = %q, email = %q, when_created = %q", selectedId, name, email, when_created)
	}
	defer db.Close()
}

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
	var id int64 = 2
	db.QueryRow("delete from accounts where id = $1", id)
	rows, err := db.Query("select * from accounts where id = $1", id)
	if err != nil {
		go_log.Logf("Error Query : %s", err.Error())
		return
	}
	var selectedId int64
	var selectedName string
	var selectedEmail string
	var when_created string
	for rows.Next() {
		rows.Scan(&selectedId, &selectedName, &selectedEmail, &when_created)
		go_log.Logf("id = %d, name = %q, email = %q, when_created = %q", selectedId, selectedName, selectedEmail, when_created)
	}
	defer db.Close()
}

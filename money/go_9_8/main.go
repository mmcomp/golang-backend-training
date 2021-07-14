package main

import (
	"database/sql"

	"github.com/lib/pq"
	go_log "github.com/mmcomp/go-log"
	"github.com/mmcomp/go-money"
)

func main() {
	connector, err := pq.NewConnector("user=mehrdad password=123456 dbname=mehrdad sslmode=disable")
	if err != nil {
		go_log.Errorf("Error Connector : %s", err.Error())
		return
	}
	db := sql.OpenDB(connector)
	defer db.Close()
	amoney := money.Cents(100)

	var name string = "test"
	var id int64
	db.QueryRow("insert into money2 (name, cents) values ($1, $2) RETURNING id", name, amoney).Scan(&id)
	go_log.Informf("The ID is %d", id)
	rows, err := db.Query("select * from money2")
	if err != nil {
		go_log.Errorf("Error Query : %s", err.Error())
		return
	}
	var selectedId int64
	var selectedName string
	var selectedCad money.CAD
	var when_created string
	for rows.Next() {
		err := rows.Scan(&selectedId, &selectedName, &selectedCad, &when_created)
		if err != nil {
			go_log.Errorf("Scan Error %s", err)
			return
		}
		go_log.Alertf("id = %d, name = %q, money = %q, when_created = %q", selectedId, selectedName, selectedCad, when_created)
	}
}

/*
CREATE TABLE money2 (
id BIGSERIAL PRIMARY KEY,
name VARCHAR ( 50 ) NOT NULL,
cents money NOT NULL,
when_created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
*/

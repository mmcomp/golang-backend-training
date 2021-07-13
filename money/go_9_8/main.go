package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
	go_log "github.com/mmcomp/go-log"
)

type CAD struct {
	cents int64
}

func Cents(n int64) CAD {
	cad := CAD{
		cents: n,
	}
	return cad
}

func ParseCAD(s string) (CAD, error) {
	s = strings.Replace(s, "$", "", 1)
	s = strings.Replace(s, "CAD", "", 1)
	s = strings.Replace(s, "¢", "", 1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, " ", "", -1)
	cad := CAD{
		cents: 0,
	}
	intValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return cad, err
	}
	cad.cents = intValue
	return cad, nil
}

func (receiver CAD) Abs() CAD {
	if receiver.cents < 0 {
		receiver.cents *= -1
	}
	return receiver
}

func (receiver CAD) Add(other CAD) CAD {
	receiver.cents += other.cents
	return receiver
}

func (receiver CAD) AsCents() int64 {
	return receiver.cents
}

func (receiver CAD) CanonicalForm() (int64, int64) {
	dollars := receiver.cents / 100
	cents := receiver.cents % 100

	return dollars, cents
}

func (receiver CAD) Mul(scalar int64) CAD {
	receiver.cents *= scalar
	return receiver
}

func (receiver CAD) Sub(other CAD) CAD {
	receiver.cents -= other.cents
	return receiver
}

func (receiver CAD) GoString() string {
	result := fmt.Sprintf("main.cents(%d)", receiver.cents)
	return result
}

func (receiver CAD) MarshalJSON() ([]byte, error) {
	result := receiver.String()
	return []byte(result), nil
}

func (receiver *CAD) UnmarshalJSON(b []byte) error {
	var cents int64
	err := json.Unmarshal(b, &cents)
	if err != nil {
		return err
	}
	receiver.cents = cents
	return nil
}

func (receiver CAD) String() string {
	var sign int8 = 1
	if receiver.cents < 0 {
		sign = -1
	}
	receiver = receiver.Abs()
	dollars, cents := receiver.CanonicalForm()
	var centsStr string = fmt.Sprintf("%d", cents)
	if cents < 10 {
		centsStr = "0" + centsStr
	}
	result := fmt.Sprintf("$%d.%s", dollars, centsStr)
	if sign < 0 {
		result = "-" + result
	}
	return result
}

func (receiver CAD) Value() (driver.Value, error) {
	return receiver.String(), nil
}

func (receiver *CAD) Scan(src interface{}) error {
	var err error
	switch casted := src.(type) {
	case string:
		*receiver, err = ParseCAD(casted)
		if err != nil {
			return err
		}
		return nil
	case []byte:
		*receiver, err = ParseCAD(string(casted))
		if err != nil {
			return err
		}
		return nil
	case int64:
		receiver.cents = casted
		return nil
	}
	return errors.New("incompatible type for CAD")
}

func main() {
	connector, err := pq.NewConnector("user=mehrdad password=123456 dbname=mehrdad sslmode=disable")
	if err != nil {
		go_log.Errorf("Error Connector : %s", err.Error())
		return
	}
	db := sql.OpenDB(connector)
	money := CAD{
		cents: 100,
	}
	var name string = "test"
	var id int64
	db.QueryRow("insert into money (name, cents) values ($1, $2)", name, money).Scan(&id)

	rows, err := db.Query("select * from money")
	if err != nil {
		go_log.Errorf("Error Query : %s", err.Error())
		return
	}
	var selectedId int64
	var selectedName string
	var selectedCad CAD
	var when_created string
	for rows.Next() {
		rows.Scan(&selectedId, &selectedName, &selectedCad, &when_created)
		go_log.Alertf("id = %d, name = %q, money = %q, when_created = %q", selectedId, selectedName, selectedCad, when_created)
	}

	defer db.Close()
}

/*
CREATE TABLE money (
id BIGSERIAL PRIMARY KEY,
name VARCHAR ( 50 ) NOT NULL,
cents integer NOT NULL,
when_created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
*/

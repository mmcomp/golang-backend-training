package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

func main() {
	money, err := ParseCAD("$1.23")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("money:", money)
}

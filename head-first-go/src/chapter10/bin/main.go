package main

import (
	"fmt"
	"headfirstgo/src/chapter10/calendar"
	"log"
)

func main() {
	//data := chapter10.Date{Year: 2019, Month: 5, Day: 25}
	date := calendar.Date{}
	err := date.SetYear(1)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("date = ", date)

	event := calendar.Event{}
	err = event.SetTitle("davirian")
	if err != nil {
		log.Fatal(err)
	}
	event.SetDay(23)
	fmt.Println(event)
}

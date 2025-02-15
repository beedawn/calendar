package main

import (
	"backend/calendar"

	"fmt"
	"time"
)


func main() {

	var cal calendar.Calendar
	var username string

	var password string
	newR := cal.NewResource()
	startTime := time.Now()
	endTime := time.Now()
	newE, _ := newR.NewEvent(startTime, endTime)
	fmt.Println(newE)

	fmt.Println("Enter username:")
	fmt.Scan(&username)
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	if username == "test" && password == "test"{
		fmt.Println("Yay!")
	}

	newR.DeleteEvent(*newE)

}

package main

import (
	"fmt"
	"backend/calendar"
)

func newLine(s string) string{
	return s + "\n"
}

func main(){

var cal calendar.Calendar
		cal.Test = "hi"

	fmt.Printf(newLine(cal.Test))
}

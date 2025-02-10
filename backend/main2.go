package main

import (
	"fmt"
	"backend/calendar"
	"backend/event"

)

func newLine(s string) string{
	return s + "\n"
}

func main(){

var cal calendar.Calendar
		cal.Test = "hi"

	fmt.Printf(newLine(cal.Test))

	cal.Room.Event = append(cal.Room.Event,event.Event{"starttime","user","duration"})


for _, event := range cal.Room.Event {
    fmt.Println(event) // Or fmt.Printf("%+v\n", event) for more detail
}
}

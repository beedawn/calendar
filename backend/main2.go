package main

import (
	"fmt"
	"backend/calendar"
	"backend/event"
	"backend/room"

)

func newLine(s string) string{
	return s + "\n"
}


func main(){

var cal calendar.Calendar
		cal.Test = "hi"

	fmt.Printf(newLine(cal.Test))
	cal.Room = make([]room.Room, 0)

	newRoom := room.Room{}
	newRoom.Event = make([]event.Event, 0)
	newRoom.Event = append(newRoom.Event, event.Event{0,"starttime","user","duration"})  
	cal.Room = append(cal.Room, newRoom)
cal.NewRoom()

cal.Room[1].NewEvent()
for i, room := range cal.Room {
	for _, event := range room.Event{

    fmt.Println(event) // Or fmt.Printf("%+v\n", event) for more detail
		}
		fmt.Println(i)
}

	cal.Room[1].DeleteEvent(event.Event{})
for i, room := range cal.Room {
	for _, event := range room.Event{

    fmt.Println(event) // Or fmt.Printf("%+v\n", event) for more detail
		}
		fmt.Printf("%d\n",room.Id)
		fmt.Println(i)
}
	fmt.Printf("%d",len(cal.Room[1].Event))
}

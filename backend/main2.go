package main

import (
	"backend/calendar"
	"backend/event"
	"backend/resource"
	"fmt"
)

func newLine(s string) string {
	return s + "\n"
}

func main() {

	var cal calendar.Calendar

	cal.Resource = make([]resource.Resource, 0)

	newResource := resource.Resource{
		Event: []event.Event{
			{Id: 0, StartTime: "starttime", User: "user", Duration: "duration"},
		},
	}
	//	newResource.Event = make([]event.Event, 0)
	//	newResource.Event = append(newResource.Event, event.Event{0, "starttime", "user", "duration"})
	cal.Resource = append(cal.Resource, newResource)
	cal.NewResource()

	cal.Resource[1].NewEvent()
	for i, resource := range cal.Resource {
		for _, event := range resource.Event {

			fmt.Println(event) // Or fmt.Printf("%+v\n", event) for more detail
		}
		fmt.Println(i)
	}

	cal.Resource[1].DeleteEvent(event.Event{})
	for i, resource := range cal.Resource {
		for _, event := range resource.Event {

			fmt.Println(event) // Or fmt.Printf("%+v\n", event) for more detail
		}
		fmt.Printf("%d\n", resource.Id)
		fmt.Println(i)
	}
	fmt.Printf("%d", len(cal.Resource[1].Event))
}

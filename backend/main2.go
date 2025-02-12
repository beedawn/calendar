package main

import (
	"backend/calendar"
	"backend/event"
	"backend/resource"
	"fmt"
	"time"
)

func newLine(s string) string {
	return s + "\n"
}

func loopEvents(c calendar.Calendar) {
	for _, resource := range c.Resource {
		for _, event := range resource.Event {

			fmt.Println(event) // Or fmt.Printf("%+v\n", event) for more detail
		}
		fmt.Printf("Resource ID: %d\n", resource.Id)
	}

}

func main() {

	var cal calendar.Calendar

	cal.Resource = make([]resource.Resource, 0)

	newResource := resource.Resource{
		Event: []event.Event{
			{Id: 0, User: "user", StartTime: time.Now(), EndTime: time.Now(), CreatedTime: time.Now()},
		},
	}
	//	newResource.Event = make([]event.Event, 0)
	//	newResource.Event = append(newResource.Event, event.Event{0, "starttime", "user", "duration"})
	cal.Resource = append(cal.Resource, newResource)
	cal.NewResource()
	startTime:=time.Now()
	endTime:=time.Now()
	cal.Resource[1].NewEvent(startTime, endTime)
	cal.Resource[1].NewEvent(startTime, time.Now())


	//loopEvents(cal)
	cal.Resource[1].DeleteEvent(event.Event{})
//	loopEvents(cal)
}

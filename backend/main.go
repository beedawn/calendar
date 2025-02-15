package main

import (
	"backend/calendar"

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

	newR := cal.NewResource()
	startTime := time.Now()
	endTime := time.Now()
	newE, _ := newR.NewEvent(startTime, endTime)

	//loopEvents(cal)
	newR.DeleteEvent(*newE)
	// loopEvents(cal)
}

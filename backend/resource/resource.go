package resource

import (
"backend/event"
	"time"
	"errors"
	"fmt"
	
)

type Resource struct {
	Id int
	Event []event.Event
	ConcurrentEvents int
	Category string
	Name string
	//what else is needed? comments?
}

func (r *Resource) conflictCheck(startTime, endTime time.Time) {
	for _, event := range r.Event{
		fmt.Println("conflict check")
fmt.Println(event)

	/*	if startTime.Before(event.EndTime){
		fmt.Println("Error! StartTime is before an Events endtime")
		}*/

		if startTime.Before(event.EndTime) && endTime.After(event.StartTime)  {
		
		fmt.Println("ahh error!")
		}
		
	}

}

func validateTimes(startTime, endTime time.Time) error{
	if endTime.Before(startTime){
		return errors.New("end date/time is before start date/time")
	}

	if endTime.Equal(startTime){
		return errors.New("end date/time is equal to start date/time")
	}
	return nil
}

func (r *Resource) NewEvent(start time.Time, end time.Time) (error) {

	if validateTimes(start, end) != nil {
		fmt.Println("error!")
	return errors.New("end time is before start time")
	}
	r.conflictCheck(start,end)
	if r.Event == nil {
	r.Event = make([]event.Event,0)
	}
	// when creating a new event in a room, we need to one, figure out how many concurrent events in a room can take place? and check to ensure there are no conflicting events before adding a new one
	//if no conflicts then
	newE := event.Event{Id:len(r.Event), CreatedTime: time.Now(), StartTime:start, EndTime: end}
	r.Event = append(r.Event, newE)
	//if there are conflicts, do nothing and return an error

	return nil
}

//for editing an event, do we just want to delete/ add event or have an editing function?

func (r *Resource) DeleteEvent(e event.Event) error {
	if r.Event == nil {
		return errors.New("Resource has no events")
	}
	for i, event := range r.Event {
		if event == e {
		r.Event = append(r.Event[:i],r.Event[i+1:]...)
			return nil
		}

	}
return errors.New("Event not found in given resouce")
}

package resource

import (
"backend/event"
	"time"
	"errors"

	
)

type Resource struct {
	Id int
	Event []event.Event
	ConcurrentEvents int
	Category string
	Name string
	//what else is needed? comments?
}

func eventOverlap (start1, end1, start2, end2 time.Time) bool {
return start1.Before(end2) && end1.After(start2)
}

func (r *Resource) conflictCheck(startTime, endTime time.Time) error{
	count := 0
	for _, event := range r.Event{
		if eventOverlap(startTime, endTime, event.StartTime, event.EndTime) {
		count++
		}
	if count >= r.ConcurrentEvents{
		return errors.New("times are conflicting!")
		}
	}
	return nil
}



func (r *Resource) NewEvent(start time.Time, end time.Time) (*event.Event, error) {


	conflictResult := r.conflictCheck(start,end)
	if conflictResult != nil{
		return nil, conflictResult
	}
	if r.Event == nil {
	r.Event = make([]event.Event,0)
	}
	// when creating a new event in a room, we need to one, figure out how many concurrent events in a room can take place? and check to ensure there are no conflicting events before adding a new one
	//if no conflicts then
	newE := event.Event{Id:len(r.Event), CreatedTime: time.Now(), StartTime:start, EndTime: end}
	r.Event = append(r.Event, newE)
	//if there are conflicts, do nothing and return an error

	return &newE, nil
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

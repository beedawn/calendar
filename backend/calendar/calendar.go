package calendar
import (
	"backend/event"
	"time"
	"errors"
	"fmt"

)

type Calendar struct {
	Id int
	Name string
	Events []event.Event
	ConcurrentEvents int

}
func eventOverlap (start1, end1, start2, end2 time.Time) bool {
return start1.Before(end2) && end1.After(start2)
}


func (c *Calendar) conflictCheck(startTime, endTime time.Time) error{
	count := 0
	for _, event := range c.Events{
		if eventOverlap(startTime, endTime, event.StartTime, event.EndTime) {
		count++
		}
	if count >= c.ConcurrentEvents && (count != 0 && c.ConcurrentEvents !=0){
		return errors.New("times are conflicting!")
		}
	}
	return nil
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

func (c *Calendar) NewEvent(start time.Time, end time.Time) (*event.Event, error) {

	if validateTimes(start, end) != nil {
		fmt.Println("error!")
	return nil, errors.New("end time is before start time")
	}
	conflictResult := c.conflictCheck(start,end)
	if conflictResult != nil{
		return nil, conflictResult
	}
	if c.Events == nil {
	c.Events = make([]event.Event,0)
	}
	// when creating a new event in a room, we need to one, figure out how many concurrent events in a room can take place? and check to ensure there are no conflicting events before adding a new one
	//if no conflicts then
	newE := event.Event{Id:len(c.Events)+1, CreatedTime: time.Now(), StartTime:start, EndTime: end}
	c.Events = append(c.Events, newE)
	//if there are conflicts, do nothing and return an error

	return &newE, nil
}

//for editing an event, do we just want to delete/ add event or have an editing function?

func (c *Calendar) DeleteEvent(e event.Event) error {
	if c.Events == nil {
		return errors.New("Resource has no events")
	}
	for i, event := range c.Events {
		if event == e {
		c.Events = append(c.Events[:i],c.Events[i+1:]...)
			return nil
		}

	}
return errors.New("Event not found in given resource")
}

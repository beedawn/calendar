package calendarmanager

import (
	"backend/calendar"
	"backend/resource"
	"backend/event"
	"fmt"
	"errors"
	"time"
)

type CalendarManager struct{
	Id int
	Calendars []calendar.Calendar
	Resources []resource.Resource

}

func (cm *CalendarManager) NewCalendar() calendar.Calendar {
	var cal calendar.Calendar

	if len (cm.Calendars) == 0 {
	fmt.Println("cal manager 0")

	cal.Id = 1
	}

	if len(cm.Calendars) != 0 {
		cal.Id = cm.Calendars[len(cm.Calendars)-1].Id+1
	}


	cm.Calendars = append (cm.Calendars, cal)

	return cal
	

}


func (cm *CalendarManager) DeleteCalendar(c calendar.Calendar) {
	num := -1
	for i, cal := range cm.Calendars {
		
		if cal.Id == c.Id{
		
		num = i+1
		}

	}

	if num != -1 {
		cm.Calendars = append(cm.Calendars[:num-1], cm.Calendars[num:]...)
	}

}




func (cm *CalendarManager) NewResource() (resource.Resource, error){
	if cm.Resources == nil {
		cm.Resources = make([]resource.Resource,0)
	}
	newR := resource.Resource{Id:len(cm.Resources)+1}
	cm.Resources = append(cm.Resources, newR)

	return newR, nil
	
}



func (cm *CalendarManager) DeleteResource(r resource.Resource) (*resource.Resource, error){
	if cm.Resources == nil {
		return nil, errors.New("No resources to delete")
	}
	for i, resource := range cm.Resources {
		if resource.Id == r.Id {
			cm.Resources = append(cm.Resources[:i],cm.Resources[i+1:]...)
			return &resource, nil
		}
	}

	return nil, errors.New("Resource not found")

}



func eventOverlap (start1, end1, start2, end2 time.Time) bool {
return start1.Before(end2) && end1.After(start2)
}


func (cm *CalendarManager) conflictCheck(startTime, endTime time.Time, r resource.Resource ) error{
	count := 0
	for _, calendar := range cm.Calendars{
	for _, event := range calendar.Events{
			for _, res := range cm.Resources{

		
		if res.Id == r.Id && eventOverlap(startTime, endTime, event.StartTime, event.EndTime) {
		count++
		}
	if count > calendar.ConcurrentEvents{
		return errors.New("times are conflicting!")
		}
			}
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

func (cm *CalendarManager) NewEvent(start time.Time, end time.Time, c *calendar.Calendar, r resource.Resource) (*event.Event, error) {

	if validateTimes(start, end) != nil {
		fmt.Println("error!")
	return nil, errors.New("end time is before start time")
	}
	conflictResult := cm.conflictCheck(start, end, r)


	if conflictResult != nil{
		return nil, conflictResult
	}
	if c.Events == nil {
	c.Events = make([]event.Event,0)
	}
	// when creating a new event in a room, we need to one, figure out how many concurrent events in a room can take place? and check to ensure there are no conflicting events before adding a new one
	//if no conflicts then
	newE := event.Event{Id: int(time.Now().UnixNano()), CreatedTime: time.Now(), StartTime:start, EndTime: end}
	c.Events = append(c.Events, newE)
	//if there are conflicts, do nothing and return an error

	return &newE, nil
}

//for editing an event, do we just want to delete/ add event or have an editing function?

func (cm *CalendarManager) DeleteEvent(e event.Event, c *calendar.Calendar,) error {
	if c.Events == nil {
		return errors.New("Resource has no events")
	}
	for i, event := range c.Events {
		if event.Id == e.Id {
		c.Events = append(c.Events[:i],c.Events[i+1:]...)
			return nil
		}

	}
return errors.New("Event not found in given resource")
}


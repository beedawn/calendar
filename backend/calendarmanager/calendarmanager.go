package calendarmanager

import (
	"backend/calendar"
	"backend/resource"
	"fmt"
	"errors"
)

type CalendarManager struct{
	Id int
	Calendars []calendar.Calendar
	Resource []resource.Resource

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
	if cm.Resource == nil {
		cm.Resource = make([]resource.Resource,0)
	}
	newR := resource.Resource{Id:len(cm.Resource)+1}
	cm.Resource = append(cm.Resource, newR)

	return newR, nil
	
}



func (cm *CalendarManager) DeleteResource(r resource.Resource) (*resource.Resource, error){
	if cm.Resource == nil {
		return nil, errors.New("No resources to delete")
	}
	for i, resource := range cm.Resource {
		if resource.Id == r.Id {
			cm.Resource = append(cm.Resource[:i],cm.Resource[i+1:]...)
			return &resource, nil
		}
	}

	return nil, errors.New("Resource not found")

}



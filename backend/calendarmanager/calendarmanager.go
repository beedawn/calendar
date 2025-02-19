package calendarmanager

import (
	"backend/calendar"
	"fmt"
)

type CalendarManager struct{
	Id int
	Calendars []calendar.Calendar

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

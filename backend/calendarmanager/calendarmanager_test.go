package calendarmanager

import (
	"testing"
)

func TestNewCalendar(t *testing.T){
	var cm CalendarManager
	cm.NewCalendar()

	if len(cm.Calendars) != 1{
		t.Errorf("Expected Calendars to have length 1, instread got %d", len(cm.Calendars))

	}	


}


func TestTwoNewCalendar(t *testing.T){
	var cm CalendarManager
	cm.NewCalendar()
	cm.NewCalendar()

	if len(cm.Calendars) != 2 {
	t.Errorf("Expected Calendars to have length 2, got %d instead", len(cm.Calendars))

	}

}


func TestDeleteCalendar(t *testing.T){
var cm CalendarManager
	cal := cm.NewCalendar()
	cm.DeleteCalendar(cal)
	if len(cm.Calendars) == 1 {
	t.Errorf("Expected Calendars to have length 1, got %d", len(cm.Calendars))
	}
}

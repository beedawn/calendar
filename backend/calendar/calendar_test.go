package calendar

import ( 
	"testing"
	"backend/event"
	"time"
)



func TestCalendarCreation (t *testing.T){
	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "user"

	newCalendar := Calendar{
		Events: []event.Event{
			{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime},
		},
	} 

	if newCalendar.Events[0].Id != 0 {
		t.Errorf("Expected Id to be 0, got %d", newCalendar.Events[0].Id)
	}
	if !newCalendar.Events[0].StartTime.Equal(startTime){
		t.Errorf("Expected startTime to be %s, got %s", startTime, newCalendar.Events[0].StartTime)
	}
	if !newCalendar.Events[0].EndTime.Equal(endTime) {
		t.Errorf("Expected endtime to be %s, got %s", endTime, newCalendar.Events[0].EndTime)
	}
	if !newCalendar.Events[0].CreatedTime.Equal(createdTime) {
		t.Errorf("Expected createdTime to be %s, got %s", createdTime, newCalendar.Events[0].CreatedTime)
	}
	if newCalendar.Events[0].User != username {
		t.Errorf("Expected user to be %s, got %s", username, newCalendar.Events[0].User)
	}
}



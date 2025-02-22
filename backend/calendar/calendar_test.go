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
		Event: []event.Event{
			{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime},
		},
	} 

	if newCalendar.Event[0].Id != 0 {
		t.Errorf("Expected Id to be 0, got %d", newCalendar.Event[0].Id)
	}
	if !newCalendar.Event[0].StartTime.Equal(startTime){
		t.Errorf("Expected startTime to be %s, got %s", startTime, newCalendar.Event[0].StartTime)
	}
	if !newCalendar.Event[0].EndTime.Equal(endTime) {
		t.Errorf("Expected endtime to be %s, got %s", endTime, newCalendar.Event[0].EndTime)
	}
	if !newCalendar.Event[0].CreatedTime.Equal(createdTime) {
		t.Errorf("Expected createdTime to be %s, got %s", createdTime, newCalendar.Event[0].CreatedTime)
	}
	if newCalendar.Event[0].User != username {
		t.Errorf("Expected user to be %s, got %s", username, newCalendar.Event[0].User)
	}
}

func TestNewEvent (t *testing.T){
	var newCalendar Calendar

	startTime := time.Now()
	endTime := time.Now()
	newCalendar.NewEvent(startTime, endTime)
	
	if len(newCalendar.Event) != 1 {
		t.Errorf("Expected length of events to be 1, got %d", len(newCalendar.Event))
	}
	
}

func TestNewEventSameTimes(t *testing.T){
	var newCalendar Calendar

	startTime := time.Now()
	endTime := time.Now()
	newCalendar.NewEvent(startTime, endTime)
	newCalendar.NewEvent(startTime, endTime)

	if len(newCalendar.Event) != 1 {
		t.Errorf("Expected length of events to be 1 when testing two events with same times, got %d", len(newCalendar.Event))
	
	}



}

func TestNewEventStartTimeAfterEndTime(t *testing.T){
	var newCalendar Calendar

	startTime := time.Now()
	endTime := time.Now()


	newCalendar.NewEvent(endTime, startTime)

	if len(newCalendar.Event) != 0 {
		t.Errorf("Expected length of events to be 0 when adding an event with a start time before the end time, got %d", len(newCalendar.Event))
	
	}



}


func TestNewEventStartTimeEqualEndTime(t *testing.T){
	var newCalendar Calendar

	startTime := time.Now()
	newCalendar.NewEvent(startTime, startTime)

	if len(newCalendar.Event) != 0 {
		t.Errorf("Expected length of events to be 0 when adding an event with a start time before the end time, got %d", len(newCalendar.Event))
	
	}



}



func TestDeleteEvent(t *testing.T){
	var newCalendar Calendar

	startTime := time.Now()
	endTime := time.Now()
	newCalendar.NewEvent(startTime, endTime)
	newCalendar.DeleteEvent(newCalendar.Event[0])

	if len(newCalendar.Event) != 0 {
		t.Errorf("Expected length of events to be 0 after deleting event, got %d", len(newCalendar.Event))
	
	}


	

}


func TestDeleteEventResourceNotExist(t *testing.T){
	var newCalendar Calendar
	newCalendar.Event = make([]event.Event,0)
	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "test"


	eventN := event.Event{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime}
		 
	err := newCalendar.DeleteEvent(eventN)
	expectedErr := "Event not found in given resource"

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}

	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}

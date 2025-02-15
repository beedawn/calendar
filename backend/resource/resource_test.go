package resource

import (
	"testing"
	"time"
	"backend/event"
)



func TestResourceCreation (t *testing.T){
	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "user"

	newResource := Resource{
		Event: []event.Event{
			{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime},
		},
	} 

	if newResource.Event[0].Id != 0 {
		t.Errorf("Expected Id to be 0, got %d", newResource.Event[0].Id)
	}
	if !newResource.Event[0].StartTime.Equal(startTime){
		t.Errorf("Expected startTime to be %s, got %s", startTime, newResource.Event[0].StartTime)
	}
	if !newResource.Event[0].EndTime.Equal(endTime) {
		t.Errorf("Expected endtime to be %s, got %s", endTime, newResource.Event[0].EndTime)
	}
	if !newResource.Event[0].CreatedTime.Equal(createdTime) {
		t.Errorf("Expected createdTime to be %s, got %s", createdTime, newResource.Event[0].CreatedTime)
	}
	if newResource.Event[0].User != username {
		t.Errorf("Expected user to be %s, got %s", username, newResource.Event[0].User)
	}
}

func TestNewEvent (t *testing.T){
	var newResource Resource

	startTime := time.Now()
	endTime := time.Now()
	newResource.NewEvent(startTime, endTime)
	
	if len(newResource.Event) != 1 {
		t.Errorf("Expected length of events to be 1, got %d", len(newResource.Event))
	}
	
}

func TestNewEventSameTimes(t *testing.T){
	var newResource Resource

	startTime := time.Now()
	endTime := time.Now()
	newResource.NewEvent(startTime, endTime)
	newResource.NewEvent(startTime, endTime)

	if len(newResource.Event) != 1 {
		t.Errorf("Expected length of events to be 1 when testing two events with same times, got %d", len(newResource.Event))
	
	}



}

func TestNewEventStartTimeAfterEndTime(t *testing.T){
	var newResource Resource

	startTime := time.Now()
	endTime := time.Now()


	newResource.NewEvent(endTime, startTime)

	if len(newResource.Event) != 0 {
		t.Errorf("Expected length of events to be 0 when adding an event with a start time before the end time, got %d", len(newResource.Event))
	
	}



}


func TestNewEventStartTimeEqualEndTime(t *testing.T){
	var newResource Resource

	startTime := time.Now()
	newResource.NewEvent(startTime, startTime)

	if len(newResource.Event) != 0 {
		t.Errorf("Expected length of events to be 0 when adding an event with a start time before the end time, got %d", len(newResource.Event))
	
	}



}



func TestDeleteEvent(t *testing.T){
	var newResource Resource

	startTime := time.Now()
	endTime := time.Now()
	newResource.NewEvent(startTime, endTime)
	newResource.DeleteEvent(newResource.Event[0])

	if len(newResource.Event) != 0 {
		t.Errorf("Expected length of events to be 0 after deleting event, got %d", len(newResource.Event))
	
	}


	

}


func TestDeleteEventResourceNotExist(t *testing.T){
	var newResource Resource
	newResource.Event = make([]event.Event,0)
	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "test"


	eventN := event.Event{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime}
		 
	err := newResource.DeleteEvent(eventN)
	expectedErr := "Event not found in given resource"

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}

	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}





func TestDeleteEventResourceEventsEmpty(t *testing.T){
	var newResource Resource

	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "test"


	eventN := event.Event{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime}
		 
	err := newResource.DeleteEvent(eventN)
	expectedErr := "Resource has no events"

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}

	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}

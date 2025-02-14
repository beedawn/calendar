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
		t.Errorf("Expected Id to be 0, got %d", newResource.Id)
	}
	if newResource.Event[0].StartTime!=startTime{
		t.Errorf("Expected startTime to be %s, got %s", startTime, newResource.Event[0].StartTime)
	}
	if newResource.Event[0].EndTime != endTime {
		t.Errorf("Expected endtime to be %s, got %s", endTime, newResource.Event[0].EndTime)
	}
	if newResource.Event[0].CreatedTime != createdTime {
		t.Errorf("Expected createdTime to be %s, got %s", createdTime, newResource.Event[0].CreatedTime)
	}
	if newResource.Event[0].User != username {
		t.Errorf("Expected user to be %s, got %s", username, newResource.Event[0].User)
	}
}

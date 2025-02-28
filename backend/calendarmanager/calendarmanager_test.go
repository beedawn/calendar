package calendarmanager

import (
	"backend/resource"
	"testing"

	"backend/event"
	"time"
)

func TestNewCalendar(t *testing.T) {
	var cm CalendarManager
	cm.NewCalendar()

	if len(cm.Calendars) != 1 {
		t.Errorf("Expected Calendars to have length 1, instread got %d", len(cm.Calendars))
	}
}

func TestTwoNewCalendar(t *testing.T) {
	var cm CalendarManager
	cm.NewCalendar()
	cm.NewCalendar()

	if len(cm.Calendars) != 2 {
		t.Errorf("Expected Calendars to have length 2, got %d instead", len(cm.Calendars))

	}

}

func TestDeleteCalendar(t *testing.T) {
	var cm CalendarManager
	cal := cm.NewCalendar()
	cm.DeleteCalendar(cal)
	if len(cm.Calendars) == 1 {
		t.Errorf("Expected Calendars to have length 1, got %d", len(cm.Calendars))
	}
}

func TestCalendarCreation(t *testing.T) {
	var cm CalendarManager

	cm.Resources = make([]resource.Resource, 0)

	if len(cm.Resources) != 0 {
		t.Errorf("Expected Resource to have a length, and for it to be 0, got %d", len(cm.Resources))
	}
}

func TestNewResource(t *testing.T) {
	var cm CalendarManager
	cm.NewResource()

	if len(cm.Resources) != 1 {
		t.Errorf("Expected Resource to have length of 1, has length %d", len(cm.Resources))
	}
}

func TestDeleteResource(t *testing.T) {
	var cm CalendarManager
	newR, _ := cm.NewResource()

	cm.DeleteResource(newR)

	if len(cm.Resources) != 0 {
		t.Errorf("Expected Resource to be length of 0, got %d", len(cm.Resources))
	}

}
func TestNilResource(t *testing.T) {

	cm := &CalendarManager{
		Resources: nil,
	}

	expectedErr := "No resources to delete"

	_, err := cm.DeleteResource(resource.Resource{})

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}

func TestResourceNotFound(t *testing.T) {

	cm := &CalendarManager{
		Resources: make([]resource.Resource, 0),
	}

	expectedErr := "Resource not found"

	_, err := cm.DeleteResource(resource.Resource{Id: 1})

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}

func eventSetup(cm *CalendarManager) {

	cm.NewCalendar()
	cm.NewResource()

}

func TestNewEvent(t *testing.T) {
	var cm CalendarManager
	eventSetup(&cm)

	startTime := time.Now()
	endTime := time.Now()
	cm.NewEvent(startTime, endTime, &cm.Calendars[0], cm.Resources[0])

	if len(cm.Calendars[0].Events) != 1 {
		t.Errorf("Expected length of events to be 1, got %d", len(cm.Calendars[0].Events))
	}

}

func TestNewEventSameTimes(t *testing.T) {
	var cm CalendarManager

	eventSetup(&cm)

	startTime := time.Now()
	endTime := time.Now()
	cm.NewEvent(startTime, endTime, &cm.Calendars[0], cm.Resources[0])
	cm.NewEvent(startTime, endTime, &cm.Calendars[0], cm.Resources[0])

	if len(cm.Calendars[0].Events) != 1 {
		t.Errorf("Expected length of events to be 1 when testing two events with same times, got %d", len(cm.Calendars[0].Events))

	}

}

func TestNewEventStartTimeAfterEndTime(t *testing.T) {
	var cm CalendarManager

	eventSetup(&cm)

	startTime := time.Now()
	endTime := time.Now()

	cm.NewEvent(endTime, startTime, &cm.Calendars[0], cm.Resources[0])

	if len(cm.Calendars[0].Events) != 0 {
		t.Errorf("Expected length of events to be 0 when adding an event with a start time before the end time, got %d", len(cm.Calendars[0].Events))

	}

}

func TestNewEventStartTimeEqualEndTime(t *testing.T) {
	var cm CalendarManager
	eventSetup(&cm)
	startTime := time.Now()
	cm.NewEvent(startTime, startTime, &cm.Calendars[0], cm.Resources[0])

	if len(cm.Calendars[0].Events) != 0 {
		t.Errorf("Expected length of events to be 0 when adding an event with a start time before the end time, got %d", len(cm.Calendars[0].Events))

	}

}

func TestDeleteEvent(t *testing.T) {
	var cm CalendarManager

	eventSetup(&cm)
	startTime := time.Now()
	endTime := time.Now()
	cm.NewEvent(startTime, endTime, &cm.Calendars[0], cm.Resources[0])
	cm.DeleteEvent(cm.Calendars[0].Events[0], &cm.Calendars[0])

	if len(cm.Calendars[0].Events) != 0 {
		t.Errorf("Expected length of events to be 0 after deleting event, got %d", len(cm.Calendars[0].Events))

	}

}

func TestDeleteEventResourceNotExist(t *testing.T) {
	var cm CalendarManager
	eventSetup(&cm)

	cm.Calendars[0].Events = make([]event.Event, 0)
	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "test"

	eventN := event.Event{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime}

	err := cm.DeleteEvent(eventN, &cm.Calendars[0])
	expectedErr := "Event not found in given resource"

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}

	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}

func TestDeleteEventNoEvents(t *testing.T) {
	var cm CalendarManager
	eventSetup(&cm)
	startTime := time.Now()
	endTime := time.Now()
	createdTime := time.Now()
	username := "test"

	eventN := event.Event{Id: 0, User: username, StartTime: startTime, EndTime: endTime, CreatedTime: createdTime}

	err := cm.DeleteEvent(eventN, &cm.Calendars[0])
	expectedErr := "Resource has no events"

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}

	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

}

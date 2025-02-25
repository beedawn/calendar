package calendarmanager

import (
	"testing"
	"backend/resource"
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





func TestCalendarCreation(t *testing.T){
	var cm CalendarManager

	cm.Resources = make([]resource.Resource, 0)


	if len(cm.Resources) != 0 {
		t.Errorf("Expected Resource to have a length, and for it to be 0, got %d", len(cm.Resources))
	}
	}


func TestNewResource(t *testing.T){
	var cm CalendarManager
	cm.NewResource()

	if len(cm.Resources)!=1 {
		t.Errorf("Expected Resource to have length of 1, has length %d", len(cm.Resources))
	}
}

func TestDeleteResource (t *testing.T){
	var cm CalendarManager
	newR, _ := cm.NewResource()

	cm.DeleteResource(newR)

	if len(cm.Resources)!=0{
		t.Errorf("Expected Resource to be length of 0, got %d",len(cm.Resources))
	}


}
func TestNilResource (t *testing.T){
	
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



func TestResourceNotFound (t *testing.T){
	
	cm := &CalendarManager{
		Resources:make([]resource.Resource,0),
	}

	expectedErr := "Resource not found"

	_, err := cm.DeleteResource(resource.Resource{Id:1})

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

	

}

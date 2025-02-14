package calendar

import (
	"testing"
	"backend/resource"
)

func TestCalendarCreation(t *testing.T){
	var cal Calendar

	cal.Resource = make([]resource.Resource, 0)


	if len(cal.Resource) != 0 {
		t.Errorf("Expected Resource to have a length, and for it to be 0, got %d", len(cal.Resource))
	}
	}


func TestNewResource(t *testing.T){
	var cal Calendar
	cal.NewResource()

	if len(cal.Resource)!=1 {
		t.Errorf("Expected Resource to have length of 1, has length %d", len(cal.Resource))
	}
}

func TestDeleteResource (t *testing.T){
	var cal Calendar
	newR := cal.NewResource()

	cal.DeleteResource(newR)

	if len(cal.Resource)!=0{
		t.Errorf("Expected Resource to be length of 0, got %d",len(cal.Resource))
	}


}
	

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
	newR, _ := cal.NewResource()

	cal.DeleteResource(newR)

	if len(cal.Resource)!=0{
		t.Errorf("Expected Resource to be length of 0, got %d",len(cal.Resource))
	}


}
func TestNilResource (t *testing.T){
	
	cal := &Calendar{
		Resource: nil,
	}

	expectedErr := "No resources to delete"

	_, err := cal.DeleteResource(resource.Resource{})

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

	

}



func TestResourceNotFound (t *testing.T){
	
	cal := &Calendar{
		Resource:make([]resource.Resource,0),
	}

	expectedErr := "Resource not found"

	_, err := cal.DeleteResource(resource.Resource{Id:1})

	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
	}

	

}

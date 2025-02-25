package resource

import ( 
	"testing"
)

func TestResourceCreation(t *testing.T){
	resource := Resource{Id:0, Category:"test", Name:"Room"}

	if resource.Id != 0 {
		t.Errorf("Expected resource to have id 0, got %d", resource.Id)

	}
	if resource.Category != "test" {
		t.Errorf("Expected resource to have category test, got %s", resource.Category)
	}
	if resource.Name != "Room" {
		t.Errorf("Expected resource to have name Room, got %s", resource.Name)
	}

}

package user

import (
	"testing"
)

func TestUserCreation(t *testing.T){
	user := User{
		Username: "testuser",
		password: "password",
		DisplayName: "Test",
	}

	if user.Username != "testuser"{
		t.Errorf("Expected Username to be testuser, got %s", user.Username)
	}
	if user.password != "password" {
		t.Errorf("Expected password to be password, got %s", user.password)
	}
	if user.DisplayName != "Test"{
		t.Errorf("Expected DisplayName to be Test, got %s", user.DisplayName)
	}

}

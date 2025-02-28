package event

import (
	"testing"
	"time"
)

func TestEventCreation(t *testing.T) {

	now := time.Now()
	startTime := now.Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	event := Event{
		Id:          1,
		User:        "testuser",
		Duration:    "1h",
		CreatedTime: now,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	if event.Id != 1 {
		t.Errorf("Expecte Id to be 1, got %d", event.Id)
	}
	if event.User != "testuser" {
		t.Errorf("Expected User to be 'testuser', got %s", event.User)
	}
	if !event.CreatedTime.Equal(now) {
		t.Errorf("Expected CreatedTime to be %v, got %v", now, event.CreatedTime)
	}
	if !event.StartTime.Equal(startTime) {
		t.Errorf("Expected StartTime to be %v, got %v", startTime, event.StartTime)

	}
	if !event.EndTime.Equal(endTime) {
		t.Errorf("Expected EndTime to be %v, got %v", endTime, event.EndTime)
	}
}

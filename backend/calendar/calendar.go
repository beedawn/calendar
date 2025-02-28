package calendar

import (
	"backend/event"
)

type Calendar struct {
	Id               int
	Name             string
	Events           []event.Event
	ConcurrentEvents int
}

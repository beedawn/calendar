package event

import (
	"backend/resource"
	"time"
)

//maybe user added to this, or UserGroup?

type Event struct {
	Id          int
	Resources   []resource.Resource
	User        string
	Duration    string
	CreatedTime time.Time
	StartTime   time.Time
	EndTime     time.Time
}

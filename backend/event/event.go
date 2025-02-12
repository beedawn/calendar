package event
import (
"time"
)
//maybe user added to this, or UserGroup?

type Event struct {
	Id int
	User string
	Duration string
	CreatedTime time.Time
	StartTime time.Time
	EndTime time.Time
}


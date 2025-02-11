package room

import (
"backend/event"
)

type Room struct {
	Id int
	Event []event.Event
}


func (r *Room) NewEvent(){
	if r.Event == nil {
	r.Event = make([]event.Event,0)
	}

	newE := event.Event{Id:len(r.Event)}
	r.Event = append(r.Event, newE)
}

func (r *Room) DeleteEvent(e event.Event) int {
	if r.Event == nil {
		return 1
	}
	for i, event := range r.Event {
		if event == e {
		r.Event = append(r.Event[:i],r.Event[i+1:]...)
			return 0
		}

	}
return 1
}

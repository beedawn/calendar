package room

import (
"backend/event"
)

type Room struct {
	Id int
	Event []event.Event
	//what else is needed? comments?
}


func (r *Room) NewEvent(){
	if r.Event == nil {
	r.Event = make([]event.Event,0)
	}
	// when creating a new event in a room, we need to one, figure out how many concurrent events in a room can take place? and check to ensure there are no conflicting events before adding a new one
	//if no conflicts then
	newE := event.Event{Id:len(r.Event)}
	r.Event = append(r.Event, newE)
	//if there are conflicts, do nothing and return an error
}

//for editing an event, do we just want to delete/ add event or have an editing function?

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

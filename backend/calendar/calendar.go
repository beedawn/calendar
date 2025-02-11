package calendar
import (
	"backend/room"
)

type Calendar struct {
	Test string
	Room []room.Room

}


func (c *Calendar) NewRoom(){
	if c.Room == nil {
		c.Room = make([]room.Room,0)
	}
	newR := room.Room{Id:len(c.Room)}
	c.Room = append(c.Room, newR)
	
}

func (c *Calendar) DeleteRoom(r room.Room) int{
	if c.Room == nil {
		return 1
	}
	for i, room := range c.Room {
		if room.Id == r.Id {
			c.Room = append(c.Room[:i],c.Room[i+1:]...)
			return 0
		}
	}

	return 1

}



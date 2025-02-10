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
 newR := room.Room{}
	c.Room = append(c.Room, newR)
	
}

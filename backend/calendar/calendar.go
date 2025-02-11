package calendar
import (
	"backend/resource"
)

type Calendar struct {
	Test string
	Resource []resource.Resource

}


func (c *Calendar) NewResource(){
	if c.Resource == nil {
		c.Resource = make([]resource.Resource,0)
	}
	newR := resource.Resource{Id:len(c.Resource)}
	c.Resource = append(c.Resource, newR)
	
}

func (c *Calendar) DeleteResource(r resource.Resource) int{
	if c.Resource == nil {
		return 1
	}
	for i, resource := range c.Resource {
		if resource.Id == r.Id {
			c.Resource = append(c.Resource[:i],c.Resource[i+1:]...)
			return 0
		}
	}

	return 1

}



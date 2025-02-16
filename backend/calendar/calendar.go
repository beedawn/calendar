package calendar
import (
	"backend/resource"
	"errors"

)

type Calendar struct {
	Id int	
	Resource []resource.Resource

}

func (c *Calendar) NewResource() resource.Resource{
	if c.Resource == nil {
		c.Resource = make([]resource.Resource,0)
	}
	newR := resource.Resource{Id:len(c.Resource)}
	c.Resource = append(c.Resource, newR)

	return newR
	
}



func (c *Calendar) DeleteResource(r resource.Resource) error{
	if c.Resource == nil {
		return errors.New("No resources to delete")
	}
	for i, resource := range c.Resource {
		if resource.Id == r.Id {
			c.Resource = append(c.Resource[:i],c.Resource[i+1:]...)
			return nil
		}
	}

	return errors.New("Resource not found")

}



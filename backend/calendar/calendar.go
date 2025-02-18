package calendar
import (
	"backend/resource"
	"errors"

)

type Calendar struct {
	Id int	
	Resource []resource.Resource

}

func (c *Calendar) NewResource() (resource.Resource, error){
	if c.Resource == nil {
		c.Resource = make([]resource.Resource,0)
	}
	newR := resource.Resource{Id:len(c.Resource)+1}
	c.Resource = append(c.Resource, newR)

	return newR, nil
	
}



func (c *Calendar) DeleteResource(r resource.Resource) (*resource.Resource, error){
	if c.Resource == nil {
		return nil, errors.New("No resources to delete")
	}
	for i, resource := range c.Resource {
		if resource.Id == r.Id {
			c.Resource = append(c.Resource[:i],c.Resource[i+1:]...)
			return &resource, nil
		}
	}

	return nil, errors.New("Resource not found")

}



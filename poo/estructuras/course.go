package main

import "fmt"
type course struct {
	Name     string
	Price   float64
	IsFree   bool
	UsersIDs []uint
	Classes  map[uint]string
}

func (c *course) ChangePrice(price float64){
	c.Price=price
}

func NewCourse(name string,price float64,isFree bool) *course{
	if price==0{
		price=30
	}

	return &course{
		Name:name, 
		Price:price,
		IsFree: isFree,
	}

}
func (c course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Printf("\n %s \n", text)
}
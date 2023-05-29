package main

import "fmt"

type Greeter interface {
	Greet()
}
 type Byer interface{
	 Bye()
 }
type Person struct {
	Name string
}


type GreeterByer interface{
	Greeter
	Byer
}

func (p Person) Bye() {
	fmt.Println("Adios!! persona", p.Name)
}
func (p Person) String() string {
	return "Hola soy persona y mi nombre"+ p.Name
}
func (p Person) Greet() {
	fmt.Println("Hola!!", p.Name)
}
type Text string

func (t Text) Greet() {
	fmt.Println("Hola soy %s!!",t)
}
func (t Text) Bye() {
	fmt.Println("Adios!! texto", t)
}
func All(gbs...GreeterByer){
	for _,gb :=range gbs{
		gb.Bye()
		gb.Greet()
	}
}
func main() {
	p:= Person{Name: "Alejandro",}
	var t Text= "Alejandro"
	All(p,t)
	
}
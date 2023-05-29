package main

import "fmt"

type Person struct {
	Name string
	Age uint
}

func NewPerson(name string, age uint) Person{
	return Person {name,age}
}
func (n Person) Greet(){
	fmt.Println("Hola!!")
}

type Human struct{
	Age uint
	Children uint
}
func NewHuman(age, children uint) Human {
	return Human {age, children}
}

type Employee struct{
	Person
	Human
	Salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee{
	return Employee{NewPerson(name,age),NewHuman(age,children),salary}
}

func (e Employee) Payroll(){
	fmt.Println(e.Salary*0.90)
}


func main(){
	e:=NewEmployee("Alejandro",30,2,1000000)
	fmt.Println(e.Human.Age)
	e.Payroll()
}
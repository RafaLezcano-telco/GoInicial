package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

//Declaracion de alias
type myAlias= course 
//definicion de tipo
type newCourse course

type newBool bool

func main(){
	c:=newCourse{name:"Go"}
	fmt.Printf("El tipo es: %T\n",c)
}
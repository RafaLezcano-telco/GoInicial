package main

import "fmt"

func main() {
	Go := Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 13, 14},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	js := Course{}
	js.Name = "Curso de js"
	js.UsersIDs = []uint{12, 34}

	Go.PrintClasses()
	Go.ChangePrice(15.5)
	fmt.Printf("%v", Go.Price)

}
package main

func main() {
	Go:= NewCourse("Go desde cero", 12.34, false)
	Go.UsersIDs = []uint{12, 13, 14}
	Go.Classes=map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		}


	
}
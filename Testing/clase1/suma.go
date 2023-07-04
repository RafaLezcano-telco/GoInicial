package clase1

func Add(a, b int) int {
	return a + b
}

func AddAcum(numbers ...int) int {
	resp := 0
	for _, k := range numbers {
		resp += k
	}
	return resp
}
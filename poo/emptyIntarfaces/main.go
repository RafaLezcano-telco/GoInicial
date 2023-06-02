package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}
func wrapper(i interface{}){
	fmt.Printf("valor:%v, tipo %T:",i,i)

	// v,ok:=i.(string)
	// if ok{
	// fmt.Println("\t%s\n",strings.ToUpper(v))
	// }
	switch v:=i.(type){
		case string: 
			fmt.Printf("\t%s\n",strings.ToUpper(v))
		case int: 
			fmt.Println(v*2)
	 default:
		fmt.Printf("valor:%v, tipo %T:",v,v)
	}
}
func main() {
	// var e exampler
	// e.x()
	wrapper(1)
	wrapper(false)
	wrapper("hola")
}
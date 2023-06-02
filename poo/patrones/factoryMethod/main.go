package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

func (p Paypal) Pay(){
	fmt.Println("Pagado con paypal")
}
type Cash struct {
}
func (c Cash) Pay(){
	fmt.Println("Pagado con Cash")
}

type CreditCard struct {
}
func (c CreditCard) Pay(){
	fmt.Println("Pagado con CreditCard")
}
func Factory(method uint) PayMethod{
	switch method{
		case 1:
				return Paypal{}
		case 2:
				return Cash{}
		case 3: 
				return CreditCard{}
		default:
				return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite metodo de pago: 1: Paypal, 2: Cash, 3: CreditCard")
	_,err:=fmt.Scanln(&method)
	if err!=nil{
		panic("Error al leer metodo de pago")
	}
	if method>3{
		panic("debe digitar 1, 2 o 3")
	}
	payMethod:=Factory(2)
	payMethod.Pay()
}
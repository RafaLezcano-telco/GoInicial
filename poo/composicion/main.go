package main

import (
	"fmt"

	"github.com/RafaLezcano-telco/GoInicial/poo/composicion/pkg/customer"
	"github.com/RafaLezcano-telco/GoInicial/poo/composicion/pkg/invoice"
	"github.com/RafaLezcano-telco/GoInicial/poo/composicion/pkg/invoiceItem"
)

func main() {
	i := invoice.New(
			"Colombia", 
			"Popayan",
		customer.New("Alejandro","Cordoba2342", "123344"),
		[]invoiceItem.Item{
			invoiceItem.New(1,"Curso Go",12.32),
			invoiceItem.New(1,"Curso c#",1.32),
			invoiceItem.New(1,"Curso js",126.4),
		},
	)

	i.SetTotal()
	fmt.Printf("%+v",i)
}
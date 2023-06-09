package invoice

import (
	"github.com/RafaLezcano-telco/GoInicial/poo/composicion/pkg/customer"
	"github.com/RafaLezcano-telco/GoInicial/poo/composicion/pkg/invoiceItem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items invoiceItem.Items
}

func New(country, city string, client customer.Customer, items invoiceItem.Items) Invoice {
	return Invoice{
		country:country, 
		city:city,  
		client:client, 
		items:items,
	}
}

//Lo ponemos de tipo puntero para actualizar el valor
func (i *Invoice) SetTotal(){
	i.total=i.items.Total()
}


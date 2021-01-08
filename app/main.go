package main

import (
	"./src/invoice"
	"fmt"
)

func main() {
	i1 := invoice.MakeRegularInvoice("123PLN", 10000, invoice.CurrencyPLN, 1)
	i2 := invoice.MakeRegularInvoice("123EUR", 10000, invoice.CurrencyEUR, 4.55)

	invoices := make([]invoice.IInvoice, 0)
	invoices = append(invoices, i1)
	invoices = append(invoices, i2)

	for i := 0; i < len(invoices); i++ {
		fmt.Println(invoices[i].GetId())
	}
}

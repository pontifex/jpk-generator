package main

import (
	"./pkg/invoice"
	"flag"
	"fmt"
)

func main() {
	filenamePtr := flag.String("filename", "invoices.csv", "Define file with invoices")
	flag.Parse()

	fmt.Println(*filenamePtr)

	pCSV := invoice.CSVParser{}

	invoices := pCSV.Parse(*filenamePtr)

	i1 := invoice.MakeRegularInvoice("123PLN", 10000, "PLN", 1)
	i2 := invoice.MakeRegularInvoice("123EUR", 10000, "EUR", 4.55)

	invoices = append(invoices, i1)
	invoices = append(invoices, i2)

	for i := 0; i < len(invoices); i++ {
		fmt.Println(invoices[i].GetId())
	}
}

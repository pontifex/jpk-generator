package main

import (
	"./pkg/invoice"
	"flag"
	"fmt"
)

func main() {
	filenamePtr := flag.String("filename", "invoices.csv", "Define file with invoices")
	flag.Parse()

	pCSV := invoice.CSVParser{}

	invoices := pCSV.Parse(*filenamePtr)

	i1 := invoice.MakeRegularInvoice("123PLN", 10000, invoice.CurrencyPLN, 1)
	i2 := invoice.MakeRegularInvoice("123EUR", 10000, invoice.CurrencyEUR, 4.55)

	invoices = append(invoices, i1)
	invoices = append(invoices, i2)

	for i := 0; i < len(invoices); i++ {
		fmt.Println(invoices[i].GetId())
	}

	gXML := invoice.XMLGenerator{}

	xmlContent := gXML.Generate(invoices)

	fmt.Println(xmlContent)
}

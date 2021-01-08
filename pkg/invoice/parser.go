package invoice

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type IParser interface {
	Parse(filename string) []IInvoice
}

type CSVParser struct {
}

func (parser CSVParser) Parse(filename string) []IInvoice {

	result := make([]IInvoice, 0)

	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	invoices := csv.NewReader(csvFile)

	i := 0
	for {
		invoiceRow, err := invoices.Read()
		if err == io.EOF {
			break
		}

		i++

		if i == 1 {
			continue
		}

		nettAmount, _ := strconv.Atoi(invoiceRow[2])
		exchangeRate, _ := strconv.ParseFloat(invoiceRow[4], 32)

		result = append(result, MakeRegularInvoice(invoiceRow[1], nettAmount, invoiceRow[3], exchangeRate))
	}

	return result
}

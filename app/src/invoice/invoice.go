package invoice

const CurrencyPLN = "PLN"
const CurrencyEUR = "EUR"

type IInvoice interface {
	GetId() string
	GetNettAmount() int
	GetCurrency() string
}

type RegularInvoice struct {
	id           string
	nettAmount   int
	currency     string
	exchangeRate float32
}

func (i RegularInvoice) GetId() string {
	return i.id
}

func (i RegularInvoice) GetNettAmount() int {
	return i.nettAmount
}

func (i RegularInvoice) GetCurrency() string {
	return i.currency
}

func MakeRegularInvoice(id string, nettAmount int, currency string, exchangeRate float32) *RegularInvoice {
	return &RegularInvoice{id, nettAmount, currency, exchangeRate}
}

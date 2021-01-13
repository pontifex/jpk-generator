package invoice

import (
	"github.com/beevik/etree"
)

type IGenerator interface {
	Generate(invoices []IInvoice) string
}

type XMLGenerator struct {
}

func (generator XMLGenerator) Generate(invoices []IInvoice) string {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	people := doc.CreateElement("People")
	people.CreateComment("These are all known people")

	jon := people.CreateElement("Person")
	jon.CreateAttr("name", "Jon")

	sally := people.CreateElement("Person")
	sally.CreateAttr("name", "Sally")

	doc.Indent(2)

	response, _ := doc.WriteToString()

	return response
}

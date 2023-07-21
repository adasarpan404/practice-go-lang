package interfaceseggregationprinciple

import "fmt"

type Printer interface {
	PrintDocument()
}

type Scanner interface {
	ScanDocument()
}

type Faxer interface {
	FaxDocument()
}

type Simple_Printer struct{}

func (sp *Simple_Printer) PrintDocument() {
	fmt.Println("Printing document......")
}

func (sp *Simple_Printer) ScanDocument() {
	fmt.Println("Scanning document......")
}

func (sp *Simple_Printer) FaxDocument() {
	fmt.Println("Faxing document......")
}

type Office_Printer struct{}

func (op *Office_Printer) PrintDocument() {
	fmt.Println("Printing document at the office......")
}

func (op *Office_Printer) ScanDocument() {
	fmt.Println("Scanning document at the office......")
}

type Home_Printer struct{}

func (hp *Home_Printer) PrintDocument() {
	fmt.Println("Printing document at the home......")
}

func (hp *Home_Printer) ScanDocument() {
	fmt.Println("Scanning document at the home......")
}

func InterfaceSeggregation() {
	officePrinter := Office_Printer{}
	simplePrinter := Simple_Printer{}
	homePrinter := Home_Printer{}

	fmt.Println("Simple Printer")
	simplePrinter.PrintDocument()
	simplePrinter.ScanDocument()
	simplePrinter.FaxDocument()

	fmt.Println("Office Printer")
	officePrinter.PrintDocument()
	officePrinter.ScanDocument()

	fmt.Println("Home Printer")
	homePrinter.PrintDocument()
	homePrinter.ScanDocument()

}

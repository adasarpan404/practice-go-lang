package interfaceseggregationprinciple

import "fmt"

type SPrinter interface {
	PrintDocument()
	ScanDocument()
	FaxDocument()
}

type SimplePrinter struct{}

func (sp *SimplePrinter) PrintDocument() {
	fmt.Println("Printing document......")
}

func (sp *SimplePrinter) ScanDocument() {
	fmt.Println("Scanning document......")
}

func (sp *SimplePrinter) FaxDocument() {
	fmt.Println("Faxing document......")
}

type OfficePrinter struct{}

func (op *OfficePrinter) PrintDocument() {
	fmt.Println("Printing office document......")
}

func (op *OfficePrinter) ScanDocument() {
	fmt.Println("Scanning office document......")
}

type HomePrinter struct{}

func (hp *HomePrinter) PrintDocument() {
	fmt.Println("Printing document at home...")
}

func (hp *HomePrinter) ScanDocument() {
	fmt.Println("Scanning document at home...")
}

func NoInterfaceSeggregation() {
	simplePrinter := SimplePrinter{}
	officePrinter := OfficePrinter{}
	homePrinter := HomePrinter{}

	fmt.Println("The simple printer")
	simplePrinter.PrintDocument()
	simplePrinter.ScanDocument()
	simplePrinter.FaxDocument()

	fmt.Println("The office printer")
	officePrinter.PrintDocument()
	officePrinter.ScanDocument()

	fmt.Println("The home printer")
	homePrinter.PrintDocument()
	homePrinter.ScanDocument()

}

package deferstatement

import "fmt"

type Product struct {
	productName string
	unitPrice   int
}

func (p Product) save() {
	fmt.Println("Ürün kaydedildi :", p.productName)
	defer Log()
}

func Log()  {
	fmt.Println("Loglandı.")
}

func DeferExampleTestFunc()  {
	p:= Product{productName: "Laptop", unitPrice: 5000}
	p = Product{productName: "Telefon", unitPrice: 2000}

	fmt.Println("İşlem başarılı.")
	fmt.Println(p.productName)
	defer p.save()
}

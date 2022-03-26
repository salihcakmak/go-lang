package structs

import "fmt"

type product struct {
	name      string
	unitPrice float64
	brand     string
}

func StructBasicTestFunc() {
	var data product
	data.name = "Car"
	data.unitPrice = 89000
	data.brand = "Volvo"
	fmt.Println(data)

	newData := product{"Simit", 2, "Simitçi Dünyası"}
	fmt.Println(newData)

	fmt.Println(product{"Computer", 4999, "Asus"})
	//fmt.Println(product{"Phone", , ""}) Hatalı kullanım. Burada eksik alan tanımlaması yapamıyoruz.

	fmt.Println(product{name: "Phone", brand: "Xiaomi"}) // Bu yöntem ile istediğimiz alanı yazdırabiliriz.
}

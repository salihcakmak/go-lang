package maps

import "fmt"

func MapTestFunc() {
	// Key - Value Değeri tutan dizi. Json benzeri. Sayı sınırı yok
	cars := make(map[string]string)

	cars["Bmw"] = "520 d"
	cars["Audi"] = "A5"
	cars["Mercedes"] = "C 250"

	fmt.Println(cars)

	//delete()
	delete(cars, "Audi")
	fmt.Println(cars)

	nameSurname := map[string]string{"scakmak" : "Salih Çakmak", "asahin": "Ahmet Talha Şahin", "iyarikan" : "İlknur Yarıkan"}
	fmt.Println(nameSurname)

}
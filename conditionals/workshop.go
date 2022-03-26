package conditionals

import "fmt"

// Görev if else kullanarak üç sayı arasından en büyük değeri bulmak
func BiggestNumberFunc() {
	var sayi1, sayi2, sayi3 int = 11, 5, 9

	var enBuyukSayi int = sayi1

	if sayi2 > enBuyukSayi {
		enBuyukSayi = sayi2
	}

	if sayi3 > enBuyukSayi {
		enBuyukSayi = sayi3
	}

	fmt.Printf("En büyük sayı : %v\n", enBuyukSayi)
}

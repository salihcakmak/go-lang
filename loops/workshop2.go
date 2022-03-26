package loops

import (
	"fmt"
)

// Görev Kullanıcının dışarıdan girdiği sayının asal olup olmadığını kontrol etmek.
// Asal sayı sadece 1'e ve kendisine tam bölünebilen sayılardır.
func IsPrimeFunc() {
	number := 0
	isPrime := true
	fmt.Println("Bir sayı giriniz.")
	fmt.Scanln(&number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}

	if isPrime {
		fmt.Println("Girilen sayı asaldır.")
	}else{
		fmt.Println("Asal Değildir.")
	}
}

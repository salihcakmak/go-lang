package loops

import "fmt"

// Görev verilen 2 sayının arkadaş sayı olup olmadığını bulmak.
// Arkadaş sayılar kendisi hariç bölenlerinin toplamı birbirine eşit olan sayılardır.
// 220 ve 284 arkadaş sayıdır. 220 yi bölen sayıların toplamı 284 eder. 284'ü bölen sayıların toplamı 220 eder.
func FriendlyNumberFunc() {
	numberOne := 0
	numberTwo := 0
	total := 0

	fmt.Println("Birinci sayıyı giriniz")
	fmt.Scanln(&numberOne)
	fmt.Println("İkinci sayıyı giriniz")
	fmt.Scanln(&numberTwo)

	for i := 1; i <= numberOne/2; i++ {
		if numberOne%i == 0 {
			total = total + i
		}
	}

	if total == numberTwo {
		fmt.Println("Bu sayılar arkadaştır.")
	} else {
		fmt.Println("Bu sayılar arkadaş değildir")
	}
}

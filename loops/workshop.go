package loops

import (
	"fmt"
)

// Görev Kullanıcıya dışarıdan sayı girişi yaptırarak buradaki tanımlı sayıyı tahmin ettirmek. Ve kaç adımda bulduğunu yazmak
// 1 - 3 : Süper | 4 - 6 : Güzel | > 6 : Fena Değil
func GuessTheNumberFunc() {
	count := 1
	status := ""
	numberInMyMind := 80
	enteredValue := 0

	fmt.Println("Bir sayı tahmin ediniz.")
	fmt.Scanln(&enteredValue)
	for enteredValue != numberInMyMind {
		if enteredValue > numberInMyMind {
			fmt.Println("Daha küçük sayı giriniz.")
		} else if enteredValue < numberInMyMind {
			fmt.Println("Daha büyük sayı giriniz.")
		}
		fmt.Scanln(&enteredValue)
		count++
	}

	if count > 0 && count <= 3 {
		status = "Süper"
	} else if count <= 6 {
		status = "Güzel"
	} else {
		status = "Fena Değil"
	}
	fmt.Println("Tebrikler Doğru Tahmin")
	fmt.Printf("Tahmin adeti : %v Başarı Durumu : %v\n", count, status)
}

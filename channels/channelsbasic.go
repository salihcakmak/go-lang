package channels

import (
	"fmt"
	"time"
)

// Tek sayıları toplayan ve toplam bilgisini kanala ileten fonksiyon
func SumOddNumber(oddNumberCn chan int) {
	total := 0

	for i := 1; i <= 9; i += 2 {
		total = total + i
		fmt.Println("Tek sayı toplama çalışıyor.")
		time.Sleep(time.Second)
	}
	oddNumberCn <- total
}

// Çift sayıları toplayan ve toplam bilgisini kanala ileten fonksiyon
func SumEvenNumber(evenNumberCn chan int) {
	total := 0

	for i := 0; i <= 10; i += 2 {
		total = total + i
		fmt.Println("Çift sayı toplama çalışıyor.")
		time.Sleep(time.Second)
	}
	evenNumberCn <- total
}

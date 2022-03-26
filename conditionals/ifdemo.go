package conditionals

import "fmt"

// Fonksiyon adı büyük harfle tanımlanırsa dışarıdan erişilebilir. Tanımlanmazsa erişilemez.
func IfDemoFunc() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900
	var durum bool = cekilmekIstenen > hesap

	if durum {
		fmt.Println("Hesaptaki para yetersiz.")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız Hazırlanıyor.")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Printf("İşlem Tamamlandı. Kalan Bakiye : %v\n", hesap)
}

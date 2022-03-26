package conditionals

import "fmt"

func IfElseDemoFunc() {
	hesap := 1000
	cekilmekIstenen := 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız Hazırlanıyor..")
		fmt.Println("Dikkat! Hesapta para kalmadı.")
	} else {
		fmt.Println("Paranız Hazırlanıyor..")
	}
}

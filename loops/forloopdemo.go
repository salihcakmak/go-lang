package loops

import "fmt"

func ForLoopDemoFunc() {
	metin := "Hello World!"
	i := 1

	// infinite loop sonsuz döngü bitme ihtimali olmayan şartlarda meydana gelir dikkat etmeliyiz.
	// for like while (for'u while gibi kullanma örneği)
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("Bitti!")
}

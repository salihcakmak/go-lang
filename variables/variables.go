package variables

import "fmt"

/* İsimlendirme Kuralları
camelCase -> Benim Tercihim
PascalCase
snake_case
UPPER_CASE
*/

func VariableTestFunc() {
	// Burada amaç değişken mantığı ve generic yapı (Tek yeri değiştir ve doğruluğundan emin ol)
	var metin string = "Go dili örnek değişken tanımlama"
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)

	// integer
	var number int = 10
	fmt.Printf("Sayı değeri : %v\n", number)
	fmt.Printf("Sayının Karesi : %v\n", number*number)

	// float default değeri float64
	var kiloGram float64 = 3.7
	fmt.Printf("Domates %v kg\n", kiloGram)

	// bool (true 1 or false 0) Bu tür sadece 2 değer alır. Doğru veya Yanlış
	var status bool = true
	fmt.Printf("Oturum Durumu : %v\n", status)

	var isim1 string = "Ahmet"
	var isim2 string = "Mehmet"

	status = isim1 == isim2 // Bu karşılaştırma sonucu doğru veya yanlış döner. == operatörü eşit mi? kontrolü yapar.
	fmt.Printf("İki string eşit mi ? > %v\n", status)

	// ! (değil) operatörü kullanımı. Durumun tersini ifade eder.
	fmt.Printf("Mevcut durum : %v - Bunun ! değili : %v\n", status, !status)

	// Kolay değişken tanımlama ve otomatik tür atama yöntemi
	dogumYili := 1997 // Burada tür değişkene verilen değer ile otomatik olarak belirlenmiştir.
	fmt.Printf("Veri tipi : %T\n", dogumYili)
	fmt.Printf("Doğum yılı değeri : %v\n", dogumYili)

}

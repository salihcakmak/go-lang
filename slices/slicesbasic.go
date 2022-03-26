package slices

import "fmt"

// Slices diziler yerine kullanılan dinamik yapıda dizi işlemleri barındıran konudur.
// make append copy fonksiyonları bulunur.
func SlicesBasicTestFunc() {
	nameArr := make([]string, 3)

	fmt.Println(nameArr)

	nameArr[0] = "Salih"
	nameArr[1] = "Hüseyin"
	nameArr[2] = "Ahmet"

	fmt.Println(nameArr)

	nameArr = append(nameArr, "Elif")
	fmt.Println(nameArr)
	fmt.Println(len(nameArr))

}
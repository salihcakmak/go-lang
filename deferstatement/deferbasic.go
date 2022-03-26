package deferstatement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı.")
}

func C(){
	fmt.Println("C fonksiyonu çalıştı.")
}

func D()  {
	fmt.Println("D fonksiyonu çalıştı.")
}

// deferler stackta birikir lifo mantığınta çalıştırılır. Amacı fonksiyon bitiminde bir fonksiyonun çalışacağını garanti etmektir.
func B() {
	defer A()
	defer C()
	fmt.Println("B fonksiyonu çalıştı.")
}
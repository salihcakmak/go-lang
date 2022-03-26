package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (customer) save() {
	fmt.Println("Eklendi")
}

func (customer) update() {
	fmt.Println("Güncellendi")
}

func StructMethodTestFunc() {
	c := customer{firstName: "Salih", lastName: "Çakmak", age: 24}
	c.save()
	c.update()
}

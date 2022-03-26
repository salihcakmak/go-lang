package structs

import "fmt"

type customer1 struct {
	firstName string
	lastName  string
	age       int
}

// Burada türünü verdiğimiz yere değişkende atarsak örneğin x içerideki verileride atamış oluruz.
func (x customer1) save1() {
	fmt.Println("Eklendi", x.firstName)
}

func (x customer1) update1() {
	fmt.Println("Güncellendi", x.lastName)
}

func StructMethodWithValueTestFunc() {
	c := customer1{firstName: "Salih", lastName: "Çakmak", age: 24}
	c.save1()
	c.update1()
}

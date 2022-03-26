package interfaces

import "fmt"

type Mortgate struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type Motor struct{
	creditPaymentTotal float64
	MotorInfo string
	rate float64
}


type CreditCalculater interface {
	Calculate() float64
}

func (m Motor) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (m Mortgate) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculater) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func InterfaceExampleTestFunc() {
	credit1 := Mortgate{creditPaymentTotal: 100000, rate: 10, address: "Ankara"}
	credit2 := Mortgate{creditPaymentTotal: 50000, rate: 10, address: "İstanbul"}
	credit3 := Car{creditPaymentTotal: 60000, rate: 10, carInfo: "Audi A5"}
	credit4 := Motor{creditPaymentTotal: 21000, rate: 17, MotorInfo: "Kawasaki"}

	credits := []CreditCalculater{credit1, credit2, credit3, credit4}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Toplam aylık ödeme tutarı :", total)
}

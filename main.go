package main

import (
	"golesson/project"
)

//fmt format kütüphanesidir.

func main() {
	/* fmt.Println("Hello World!")
	fmt.Println("Selam!")
	variables.VariableTestFunc()
	conditionals.IfDemoFunc()
	conditionals.IfElseDemoFunc()
	conditionals.BiggestNumberFunc()
	loops.ForLoopDemoFunc()
	loops.ForLoopBasicDemoFunc()
	loops.GuessTheNumberFunc()
	loops.IsPrimeFunc()
	loops.FriendlyNumberFunc()
	arrays.IntArrayTestFunc()
	arrays.StringArrayTestFunc()
	arrays.BiggestArrayNumberFunc()
	arrays.MultidimensionalArrayTest()
	slices.SlicesBasicTestFunc()
	slices.SlicesAdvancedTestFunc()
	functions.SayHello("Salih")
	result := functions.AdditionOperation(50, 55)
	fmt.Println(result)
	// Go dilinde var olan ama kullanılmak istenmeyen değerler _ ile yok sayılabilir.
	toplamaSonuc, cikarmaSonuc, _, bolmeSonuc := functions.BasicOperations(10, 5)
	fmt.Println("Toplam :", toplamaSonuc)
	fmt.Println("Çıkarma :", cikarmaSonuc)
	//fmt.Println("Çarpma :", carpmaSonuc)
	fmt.Println("Bölme :", bolmeSonuc)
	fmt.Println(functions.GetSum(1,51,325,34,2))
	fmt.Println(functions.GetSum(1,4))
	fmt.Println(functions.GetSum())
	maps.MapTestFunc()
	rangecommand.RangeTestFunc()
	rangecommand.TotalEvenNumbers()
	rangecommand.MapWithRangeTestFunc()
	var number int = 20
	fmt.Println("Value in main before use", number)
	pointers.PointerBasicTestFunc(&number)
	fmt.Println("Value in main after use :", number)
	numbers := []int{1, 7, 12, 66, 3, 5, 10}
	fmt.Println("Array in main before use :", numbers)
	pointers.ArrayPointerFunc(numbers)
	fmt.Println("Array in main after use :", numbers)
	fmt.Println("Variable address in main :", &number)
	pointers.PointerAddressFunc(&number)
	structs.StructBasicTestFunc()
	structs.StructMethodTestFunc()
	structs.StructMethodWithValueTestFunc()
	go goroutines.WriteOddNumbers()  // go komutu işlemin ayrı bir thread de çalışacağını belirtir.
	go goroutines.WriteEvenNumbers() // fakat çok hızlı olacağı için farkı anlamak zor olabilir.
	// aynı zamanda asıl çalıştırılan program main olduğu için ve main kodları go ile belirtilmiş olan fonksiyonlar
	// çalışırken bile çalışmaya devam edip ilerlediği için kendi işi bitince kapanacaktır ve fonksiyonların işini tamamlayıp
	// tamamlamadığına bakmayacaktır. Bu yüzden bu örnek için saniye ayarı iyi verilmelidir.
	time.Sleep(time.Second * 2)
	fmt.Println("Main ended 2 second later")

	// Fonksiyonda kullanabilmek için int değer taşıyan iki kanal oluşturduk.
	tekSayiCn := make(chan int)
	ciftSayiCn := make(chan int)

	go channels.SumOddNumber(tekSayiCn)
	go channels.SumEvenNumber(ciftSayiCn)

	// Kanaldaki değeri değişkene aktardık.
	var tekSayiToplam int = <-tekSayiCn
	var ciftSayiToplam int = <-ciftSayiCn

	fmt.Println("Tek sayı toplam değeri :", tekSayiToplam)
	fmt.Println("Cift sayı toplam değeri :", ciftSayiToplam)
	fmt.Println("İki toplam değerin çarpımı :", tekSayiToplam*ciftSayiToplam)
	interfaces.InterfaceBasicTestFunc()
	interfaces.InterfaceExampleTestFunc()

	deferstatement.B()
	deferstatement.DeferDemoTestFunc()
	deferstatement.DeferExampleTestFunc()

	errorhandling.ErrorHandlingTestFunc()
	interfaces.TypeAssertionTestFunc()
	str, err := errorhandling.GuessFuncWithError(111)
	fmt.Println(str)
	fmt.Println(err)

	str, err = errorhandling.GuessFuncWithSelfError(121)
	fmt.Println(str)
	fmt.Println(err)
	stringfunctions.StringFunctionTest()
	restful.RestApiGetFunc()
	restful.RestApiPostFunc() */
	//project.GetAllProducts()
	project.AddProduct()
}

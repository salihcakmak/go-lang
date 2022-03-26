package errorhandling

import (
	"fmt"
	"os"
)
// type assertion and error handling
func ErrorHandlingTestFunc() {
	file, err := os.Open("Demo.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya Bulunamadı :", pErr.Path)
			return
		}else {
			fmt.Println("Dosya Hatası")
			return
		}
	}else {
		fmt.Println(file.Name())
	}
}

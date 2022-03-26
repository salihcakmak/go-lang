package errorhandling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

func (b borderException) Error() string {
	return fmt.Sprintf("%d----%s", b.parameter, b.message)
}

func GuessFuncWithSelfError(guessNumber int) (string, error) {
	if guessNumber < 1 || guessNumber > 100 {
		return "", &borderException{guessNumber, "Sınırların dışında kaldı"}
	} else {
		return "Doğru değer aralığı", nil
	}
}

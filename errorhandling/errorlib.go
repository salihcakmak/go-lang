package errorhandling

import "errors"

func GuessFuncWithError(guessNumber int) (string, error) {
	numberInMyMind := 50

	if guessNumber < 1 || guessNumber > 100 {
		return "", errors.New("1-100 Arasında bir değer giriniz")
	} else if guessNumber > numberInMyMind {
		return "Daha küçük sayı giriniz", nil
	} else if guessNumber < numberInMyMind {
		return "Daha büyük sayı giriniz", nil
	} else {
		return "Bildiniz", nil
	}
}

package generators

import (
	"math/rand"
	"strconv" // Import strconv
)

func GenerateGiftcardNumber(length int, prefix string) string {
	baseLength := length - len(prefix) - 1 // -1 para el dígito de verificación
	base := make([]byte, baseLength)

	for i := 0; i < baseLength; i++ {
		base[i] = byte('0' + rand.Intn(10))
	}

	number := prefix + string(base)
	checkDigit := GenerateCheckDigit(number)

	return number + strconv.Itoa(checkDigit)
}

func GenerateCheckDigit(number string) int {
	sum := 0
	for i, digit := range number {
		val := int(digit - '0')
		if i%2 == 0 {
			val *= 2
		}
		sum += val / 10
		sum += val % 10
	}
	return (10 - (sum % 10)) % 10
}

package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
)

// ErrInvalidString - ошибка распаковки
var ErrInvalidString = errors.New("invalid string")

// Unpack - Распаковка строки
func Unpack(str string) (string, error) {

	// разделяем строку
	splitStr := strings.Split(str, "")

	// проверяем строку len = 0
	if len(splitStr) == 0 {
		return "", nil
	}

	// проверяем строку len = 1
	if len(splitStr) == 1 {

		// проверяем на число первый символ
		if strings.Contains("0123456789", splitStr[0]) {
			return "", ErrInvalidString
		}

		return splitStr[0], nil
	}

	// строка len = 2 и больше
	for i := 0; i < len(splitStr)-1; i++ {

		// пустую строку пропускаем
		if splitStr[i] == "" {
			continue
		}

		// если число то ошибка
		if strings.Contains("0123456789", splitStr[i]) {
			return "", ErrInvalidString
		}

		// если следующее - число
		if strings.Contains("0123456789", splitStr[i+1]) {

			repeatCount := strings.Index("0123456789", splitStr[i+1])
			repeatStr := strings.Repeat(splitStr[i], repeatCount)

			splitStr[i] = repeatStr
			splitStr[i+1] = ""
		}

	}

	// соединяем строку
	resStr := strings.Join(splitStr[:], "")

	return resStr, nil
}

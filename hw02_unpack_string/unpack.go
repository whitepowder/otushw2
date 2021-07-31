package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	strBuilder       strings.Builder
	runeArray        []rune
	splitedStr       []string
)

func Unpack(inputStr string) (string, error) {
	strBuilder.Reset()
	runeArray = []rune(inputStr)
	for i := 0; i < len(runeArray); i++ {
		switch {
		case unicode.IsDigit(runeArray[i]) && i == 0:
			return "", ErrInvalidString
		case unicode.IsDigit(runeArray[i]) && unicode.IsDigit(runeArray[i-1]):
			return "", ErrInvalidString
		default:
			BuiltString(inputStr, i)
		}
	}
	return strBuilder.String(), nil
}

func getRepeatValue(str string, position int) (int, error) {
	splitedStr = strings.Split(str, "")
	switch {
	case unicode.IsDigit(runeArray[position]):
		return strconv.Atoi(splitedStr[position])
	default:
		return 1, nil
	}
}

func BuiltString(str string, position int) {
	repeatValue, _ := getRepeatValue(str, position)
	switch {
	case position != 0 && !unicode.IsDigit(runeArray[position-1]):
		strBuilder.WriteString(strings.Repeat(splitedStr[position-1], repeatValue))
		if (position+1 == len(runeArray)) && !unicode.IsDigit(runeArray[position]) {
			strBuilder.WriteString(strings.Repeat(splitedStr[position], repeatValue))
		}
	case (position+1 == len(runeArray)) && !unicode.IsDigit(runeArray[position]):
		strBuilder.WriteString(strings.Repeat(splitedStr[position], repeatValue))
	}
}

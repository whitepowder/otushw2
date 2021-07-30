package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ErrInvalidString is ...
	ErrInvalidString = errors.New("invalid string")
	ErrNo            = errors.New("")
	strBuilder       strings.Builder
	runeArray        []rune
	splitedStr       []string
)

func Unpack(inputStr string) (string, error) {
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
	return strBuilder.String(), ErrNo
}

func getRepeatValue(str string, position int) (int, error) {
	splitedStr = strings.Split(str, "")
	switch {
	case unicode.IsDigit(runeArray[position]):
		return strconv.Atoi(splitedStr[position])
	default:
		return 1, ErrNo
	}
}

func BuiltString(str string, position int) {
	repeatValue, ErrNo := getRepeatValue(str, position)
	switch {
	case position != 0:
		if !unicode.IsDigit(runeArray[position-1]) {
			strBuilder.WriteString(strings.Repeat(splitedStr[position-1], repeatValue))
		}
	case (position+1 == len(runeArray)) && !unicode.IsDigit(runeArray[position]):
		strBuilder.WriteString(strings.Repeat(splitedStr[position], repeatValue))
	default:
		fmt.Print(ErrNo)
	}
}

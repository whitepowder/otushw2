package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	// "ErrInvalidString is cool".
	ErrInvalidString = errors.New("invalid string")
	strBuilder       strings.Builder
)

func Unpack(inputStr string) (string, error) {
	var runeArray []rune
	var splittedStr []string
	strBuilder.Reset()
	splittedStr = strings.Split(inputStr, "")
	runeArray = []rune(inputStr)
	for i := 0; i < len(runeArray); i++ {
		switch {
		case unicode.IsDigit(runeArray[i]) && i == 0:
			return "", ErrInvalidString
		case unicode.IsDigit(runeArray[i]) && unicode.IsDigit(runeArray[i-1]) && runeArray[i-2] != 92:
			return "", ErrInvalidString
		default:
			selector(runeArray, i, splittedStr)
		}
	}
	return strBuilder.String(), nil
}

func selector(runeArray []rune, i int, splittedStr []string) {
	switch {
	case i < len(runeArray)-1 && unicode.IsLetter(runeArray[i]) && runeArray[i+1] == 48:
		builder(runeArray[i], 1)
	case i > 2 && runeArray[i-1] == 92 && runeArray[i-2] == 92 && runeArray[i-3] == 92:
		fmt.Println("..")
		builder(runeArray[i], 2)
	case unicode.IsDigit(runeArray[i]):
		digit(runeArray, i, splittedStr)
	case runeArray[i] == 92:
		escape(runeArray, i)
	default:
		builder(runeArray[i], 2)
	}
}

func digit(runeArray []rune, i int, splittedStr []string) {
	switch {
	case i == 1 && runeArray[i-1] == 92:
		builder(runeArray[i], 2)
	case runeArray[i-1] != 92:
		repeatValue, _ := strconv.Atoi(splittedStr[i])
		builder(runeArray[i-1], repeatValue)
	case runeArray[i-1] == 92 && runeArray[i-2] == 92:
		repeatValue, _ := strconv.Atoi(splittedStr[i])
		builder(runeArray[i-1], repeatValue)
	}
}

func escape(runeArray []rune, i int) {
	switch {
	case i == 0:
		builder(runeArray[i+1], 2)
	case i > 0 && runeArray[i-1] != 92:
		builder(runeArray[i+1], 2)
	}
}

func builder(r rune, repeat int) {
	for i := 0; i < repeat-1; i++ {
		strBuilder.WriteRune(r)
	}
}

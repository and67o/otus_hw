package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var resStr strings.Builder

	firstChar, _ := utf8.DecodeRuneInString(str)

	if unicode.IsDigit(firstChar) || haveDoubleNumbers(str) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(str); i++ {
		value := str[i]
		if !unicode.IsDigit(rune(value)) {
			nextIndex := i + 1
			if nextIndex == len(str) {
				nextIndex = i
			}
			if nextIndex < len(str) {
				nexValue := str[nextIndex]
				if unicode.IsDigit(rune(nexValue)) {
					countOfAdd, _ := strconv.Atoi(string(nexValue))
					strAdd := strings.Repeat(string(value), countOfAdd)
					resStr.WriteString(strAdd)
				} else {
					resStr.WriteString(string(value))
				}
			}
		}
	}
	return resStr.String(), nil
}

func haveDoubleNumbers(str string) bool {
	reg := regexp.MustCompile("\\d{2,}")
	doubleNumber := reg.FindString(str)

	return doubleNumber != ""
}

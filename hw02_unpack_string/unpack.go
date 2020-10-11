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
	if str == "" {
		return "", nil
	}

	firstChar, _ := utf8.DecodeRuneInString(str)

	if unicode.IsDigit(firstChar) || haveDoubleNumbers(str) {
		return "", ErrInvalidString
	}

	var resStr strings.Builder
	var prevChar string

	for _, v := range str {
		currentChar := string(v)
		if unicode.IsDigit(v) {
			count, err := strconv.Atoi(currentChar)
			if err != nil {
				return "", err
			}

			if count > 0 {
				strAdd := strings.Repeat(prevChar, count-1)
				resStr.WriteString(strAdd)
			} else {
				resStrTmp := resStr.String()
				resStr.Reset()
				resStrTmp = strings.TrimSuffix(resStrTmp, prevChar)
				resStr.WriteString(resStrTmp)
			}
		} else {
			resStr.WriteString(currentChar)
		}
		prevChar = currentChar
	}

	return resStr.String(), nil
}

func haveDoubleNumbers(str string) bool {
	reg := regexp.MustCompile(`\d{2,}`)
	doubleNumber := reg.FindString(str)

	return doubleNumber != ""
}

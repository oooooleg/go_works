package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	sb := &strings.Builder{}
	var curRune rune
	var escaped bool
	var err error

	for _, r := range s {
		curRune, escaped, err = processRune(r, escaped, curRune, sb)
		if err != nil {
			return "", err
		}
	}

	if curRune != 0 {
		sb.WriteRune(curRune)
	}
	return sb.String(), nil
}

func processRune(r rune, escaped bool, curRune rune, sb *strings.Builder) (rune, bool, error) {
	if r == '\\' {
		if escaped {
			return r, false, nil
		}
		if curRune != 0 {
			sb.WriteRune(curRune)
		}
		return 0, true, nil
	}

	if curRune == 0 {
		if unicode.IsDigit(r) {
			if !escaped {
				return 0, false, ErrInvalidString
			}
		}
		return r, false, nil
	}

	if !unicode.IsDigit(r) || escaped {
		sb.WriteRune(curRune)
		return r, false, nil
	}

	count := (int)(r - '0')

	if count > 9 || count < 0 {
		return 0, false, ErrInvalidString
	}

	for i := 0; i < count; i++ {
		sb.WriteRune(curRune)
	}

	return 0, false, nil
}

package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func isFLetter(s string) error {
	if !unicode.IsLetter(rune(s[0])) {
		return ErrInvalidString
	}
	return nil
}

func isOneNum(s rune) error {
	if unicode.IsDigit(s) {
		return ErrInvalidString
	}
	return nil
}

func Unpack(s string) (string, error) {
	// If the curr length is empty, just return
	if len(s) < 1 {
		return s, nil
	}
	err := isFLetter(s)
	// OMg how to handle more compact the errors ?
	if err != nil {
		return "", err
	}

	var (
		expStr strings.Builder
		accum  = ""
		dig    int
	)

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			err := isOneNum(rune(s[i-1]))
			if err != nil {
				return "", err
			}

			dig, _ = strconv.Atoi(string(s[i]))
			if dig != 0 {
				expStr.WriteString(strings.Repeat(string(s[i-1]), dig-1))
			} else {
				// Flush previous val, if the current number is 0
				accum = ""
			}
		} else {
			if len(accum) > 0 {
				expStr.WriteString(accum)
			}
			accum = string(s[i])
		}
	}

	expStr.WriteString(accum)
	return expStr.String(), err
}

func main() {
	str := `a\b2s2d3j`
	res, err := Unpack(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

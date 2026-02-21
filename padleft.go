package main

import "strings"

// Pad takes a string and an int, and pads the string
// So it is at least N characters long, using whitespace on
// The left hand side to make up the extra length
// It returns an error in cases where n < 0
func Pad(s string, n int) (string, error) {
	return PadCharacter(s, n, ' ')
}

// PadCharacter takes a string, an int, and a rune, and pads the string
// So it is at least N characters long, using whitespace on
// The left hand side to make up the extra length
// It returns an error in cases where n < 0
func PadCharacter(s string, count int, char rune) (string, error) {
	if count < 0 {
		return "", ErrInvalidPadLength
	}
	stringSize := len(s)
	if stringSize > count {
		// Return just the string
		return s, nil
	}
	newString := strings.Repeat(string(char), count-stringSize) + s
	return newString, nil
}

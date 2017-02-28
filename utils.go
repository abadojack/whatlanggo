package whatlanggo

import "unicode"

//isStopChar returns true if r is space, punctuation or digit.
func isStopChar(r rune) bool {
	if unicode.IsSymbol(r) || unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

//abs returns the absolute value of x.
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

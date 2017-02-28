package whatlanggo

import "unicode"

//Info represents a full outcome of language detection.
type Info struct {
	Lang   Lang
	Script *unicode.RangeTable
}

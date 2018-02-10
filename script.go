package whatlanggo

import "unicode"

type scriptCounter struct {
	checkFunc func(r rune) bool
	script    *unicode.RangeTable
	count     *int
}

//Scripts is the set of Unicode script tables.
var Scripts = map[*unicode.RangeTable]string{
	unicode.Latin:      "Latin",
	unicode.Cyrillic:   "Cyrillic",
	unicode.Arabic:     "Arabic",
	unicode.Devanagari: "Devanagari",
	unicode.Hiragana:   "Hiragana",
	unicode.Katakana:   "Katakana",
	unicode.Ethiopic:   "Ethiopic",
	unicode.Hebrew:     "Hebrew",
	unicode.Bengali:    "Bengali",
	unicode.Georgian:   "Georgian",
	unicode.Han:        "Han",
	unicode.Hangul:     "Hangul",
	unicode.Greek:      "Greek",
	unicode.Kannada:    "Kannada",
	unicode.Tamil:      "Tamil",
	unicode.Thai:       "Thai",
	unicode.Gujarati:   "Gujarati",
	unicode.Gurmukhi:   "Gurmukhi",
	unicode.Telugu:     "Telugu",
	unicode.Malayalam:  "Malayalam",
	unicode.Oriya:      "Oriya",
	unicode.Myanmar:    "Myanmar",
	unicode.Sinhala:    "Sinhala",
	unicode.Khmer:      "Khmer",
}

//DetectScript returns only the script of the given text.
func DetectScript(text string) *unicode.RangeTable {
	halfLen := len(text) / 2

	scriptCounter := []scriptCounter{
		{isLatin, unicode.Latin, new(int)},
		{isCyrillic, unicode.Cyrillic, new(int)},
		{isArabic, unicode.Arabic, new(int)},
		{isDevanagari, unicode.Devanagari, new(int)},
		{isHiraganaKatakana, _HiraganaKatakana, new(int)},
		{isEthiopic, unicode.Ethiopic, new(int)},
		{isHebrew, unicode.Hebrew, new(int)},
		{isBengali, unicode.Bengali, new(int)},
		{isGeorgian, unicode.Georgian, new(int)},
		{isHan, unicode.Han, new(int)},
		{isHangul, unicode.Hangul, new(int)},
		{isGreek, unicode.Greek, new(int)},
		{isKannada, unicode.Kannada, new(int)},
		{isTamil, unicode.Tamil, new(int)},
		{isThai, unicode.Thai, new(int)},
		{isGujarati, unicode.Gujarati, new(int)},
		{isGurmukhi, unicode.Gurmukhi, new(int)},
		{isTelugu, unicode.Telugu, new(int)},
		{isMalayalam, unicode.Malayalam, new(int)},
		{isOriya, unicode.Oriya, new(int)},
		{isMyanmar, unicode.Myanmar, new(int)},
		{isSinhala, unicode.Sinhala, new(int)},
		{isKhmer, unicode.Khmer, new(int)},
	}

	for _, ch := range text {
		if isStopChar(ch) {
			continue
		}

		for i, sc := range scriptCounter {
			if sc.checkFunc(ch) {
				*sc.count++
				if *sc.count > halfLen {
					return sc.script
				}

				//if script is found, move it closer to the front so that it be checked first.
				if i > 0 {
					scriptCounter[i], scriptCounter[i-1] = scriptCounter[i-1], scriptCounter[i]
				}
			}
		}
	}

	//find the script that occurs the most in the text and return it.
	jpCount := 0
	max := 0
	maxScript := &unicode.RangeTable{}
	for _, script := range scriptCounter {
		if *script.count > max {
			max = *script.count
			maxScript = script.script
			if script.script == _HiraganaKatakana {
				jpCount = max
			}
		}
	}

	switch {
	case max == 0:
		//if no valid script is detected, return nil.
		return nil
	case max != 0 && (maxScript == unicode.Han && jpCount > 0):
		// If Hiragana or Katakana is included, even if judged as Mandarin,
		// it is regarded as Japanese. Japanese uses Kanji (unicode.Han)
		// in addition to Hiragana and Katakana.
		return _HiraganaKatakana
	default:
		return maxScript
	}
}

var isCyrillic = func(r rune) bool {
	return unicode.Is(unicode.Cyrillic, r)
}

var isLatin = func(r rune) bool {
	return unicode.Is(unicode.Latin, r)
}

var isArabic = func(r rune) bool {
	return unicode.Is(unicode.Arabic, r)
}

var isDevanagari = func(r rune) bool {
	return unicode.Is(unicode.Devanagari, r)
}

var isEthiopic = func(r rune) bool {
	return unicode.Is(unicode.Ethiopic, r)
}

var isHebrew = func(r rune) bool {
	return unicode.Is(unicode.Hebrew, r)
}

var isHan = func(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

var isBengali = func(r rune) bool {
	return unicode.Is(unicode.Bengali, r)
}

var isHiraganaKatakana = func(r rune) bool {
	return unicode.Is(_HiraganaKatakana, r)
}

var isHangul = func(r rune) bool {
	return unicode.Is(unicode.Hangul, r)
}

var isGreek = func(r rune) bool {
	return unicode.Is(unicode.Greek, r)
}

var isKannada = func(r rune) bool {
	return unicode.Is(unicode.Kannada, r)
}

var isTamil = func(r rune) bool {
	return unicode.Is(unicode.Tamil, r)
}

var isThai = func(r rune) bool {
	return unicode.Is(unicode.Thai, r)
}

var isGujarati = func(r rune) bool {
	return unicode.Is(unicode.Gujarati, r)
}

var isGurmukhi = func(r rune) bool {
	return unicode.Is(unicode.Gurmukhi, r)
}

var isTelugu = func(r rune) bool {
	return unicode.Is(unicode.Telugu, r)
}

var isMalayalam = func(r rune) bool {
	return unicode.Is(unicode.Malayalam, r)
}

var isOriya = func(r rune) bool {
	return unicode.Is(unicode.Oriya, r)
}

var isMyanmar = func(r rune) bool {
	return unicode.Is(unicode.Myanmar, r)
}

var isSinhala = func(r rune) bool {
	return unicode.Is(unicode.Sinhala, r)
}

var isKhmer = func(r rune) bool {
	return unicode.Is(unicode.Khmer, r)
}

var isGeorgian = func(r rune) bool {
	return unicode.Is(unicode.Georgian, r)
}

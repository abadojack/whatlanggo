package whatlanggo

import "unicode"

type scriptCounter struct {
	checkFunc func(r rune) bool
	script    *unicode.RangeTable
	count     int
}

// Scripts is the set of Unicode script tables.
var Scripts = map[*unicode.RangeTable]string{
	unicode.Arabic:     "Arabic",
	unicode.Bengali:    "Bengali",
	unicode.Cyrillic:   "Cyrillic",
	unicode.Ethiopic:   "Ethiopic",
	unicode.Devanagari: "Devanagari",
	unicode.Han:        "Han",
	unicode.Georgian:   "Georgian",
	unicode.Greek:      "Greek",
	unicode.Gujarati:   "Gujarati",
	unicode.Gurmukhi:   "Gurmukhi",
	unicode.Hangul:     "Hangul",
	unicode.Hebrew:     "Hebrew",
	unicode.Hiragana:   "Hiragana",
	unicode.Kannada:    "Kannada",
	unicode.Katakana:   "Katakana",
	unicode.Khmer:      "Khmer",
	unicode.Latin:      "Latin",
	unicode.Malayalam:  "Malayalam",
	unicode.Myanmar:    "Myanmar",
	unicode.Oriya:      "Oriya",
	unicode.Sinhala:    "Sinhala",
	unicode.Tamil:      "Tamil",
	unicode.Telugu:     "Telugu",
	unicode.Thai:       "Thai",
}

// DetectScript returns only the script of the given text.
func DetectScript(text string) *unicode.RangeTable {
	halfLen := len(text) / 2

	scriptCounter := []scriptCounter{
		{isLatin, unicode.Latin, 0},
		{isCyrillic, unicode.Cyrillic, 0},
		{isArabic, unicode.Arabic, 0},
		{isDevanagari, unicode.Devanagari, 0},
		{isHiraganaKatakana, _HiraganaKatakana, 0},
		{isEthiopic, unicode.Ethiopic, 0},
		{isHebrew, unicode.Hebrew, 0},
		{isBengali, unicode.Bengali, 0},
		{isGeorgian, unicode.Georgian, 0},
		{isHan, unicode.Han, 0},
		{isHangul, unicode.Hangul, 0},
		{isGreek, unicode.Greek, 0},
		{isKannada, unicode.Kannada, 0},
		{isTamil, unicode.Tamil, 0},
		{isThai, unicode.Thai, 0},
		{isGujarati, unicode.Gujarati, 0},
		{isGurmukhi, unicode.Gurmukhi, 0},
		{isTelugu, unicode.Telugu, 0},
		{isMalayalam, unicode.Malayalam, 0},
		{isOriya, unicode.Oriya, 0},
		{isMyanmar, unicode.Myanmar, 0},
		{isSinhala, unicode.Sinhala, 0},
		{isKhmer, unicode.Khmer, 0},
	}

	for _, ch := range text {
		if isStopChar(ch) {
			continue
		}

		for i, sc := range scriptCounter {
			if sc.checkFunc(ch) {
				scriptCounter[i].count++
				if scriptCounter[i].count > halfLen {
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
		if script.count > max {
			max = script.count
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

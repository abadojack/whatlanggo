package whatlanggo

import (
	"sort"
	"unicode"
)

// Detect language and script of the given text.
func Detect(text string) Info {
	return DetectWithOptions(text, Options{})
}

// DetectLang detects only the language by a given text.
func DetectLang(text string) Lang {
	return Detect(text).Lang
}

// DetectLangWithOptions detects only the language of the given text with the provided options.
func DetectLangWithOptions(text string, options Options) Lang {
	return DetectWithOptions(text, options).Lang
}

// DetectWithOptions detects the language and script of the given text with the provided options.
func DetectWithOptions(text string, options Options) Info {
	script := DetectScript(text)
	if script != nil {
		lang, confidence := detectLangBaseOnScript(text, options, script)
		return Info{
			Lang:       lang,
			Script:     script,
			Confidence: confidence,
		}
	}

	return Info{
		Lang:       -1,
		Script:     nil,
		Confidence: 0,
	}
}

func detectLangBaseOnScript(text string, options Options, script *unicode.RangeTable) (Lang, float64) {
	switch script {
	case unicode.Latin:
		return detectLangInProfiles(text, options, latinLangs)
	case unicode.Cyrillic:
		return detectLangInProfiles(text, options, cyrillicLangs)
	case unicode.Devanagari:
		return detectLangInProfiles(text, options, devanagariLangs)
	case unicode.Hebrew:
		return detectLangInProfiles(text, options, hebrewLangs)
	case unicode.Ethiopic:
		return detectLangInProfiles(text, options, ethiopicLangs)
	case unicode.Arabic:
		return detectLangInProfiles(text, options, arabicLangs)
	case unicode.Han:
		return Cmn, 1
	case unicode.Bengali:
		return Ben, 1
	case unicode.Hangul:
		return Kor, 1
	case unicode.Georgian:
		return Kat, 1
	case unicode.Greek:
		return Ell, 1
	case unicode.Kannada:
		return Kan, 1
	case unicode.Tamil:
		return Tam, 1
	case unicode.Thai:
		return Tha, 1
	case unicode.Gujarati:
		return Guj, 1
	case unicode.Gurmukhi:
		return Pan, 1
	case unicode.Telugu:
		return Tel, 1
	case unicode.Malayalam:
		return Mal, 1
	case unicode.Oriya:
		return Ori, 1
	case unicode.Myanmar:
		return Mya, 1
	case unicode.Sinhala:
		return Sin, 1
	case unicode.Khmer:
		return Khm, 1
	case _HiraganaKatakana:
		return Jpn, 1
	default:
		return -1, 0
	}
}

type langDistance struct {
	lang Lang
	dist int
}

func detectLangInProfiles(text string, options Options, langProfileList langProfileList) (Lang, float64) {
	trigrams := getTrigramsWithPositions(text)

	langDistances := []langDistance{}

	for lang, langTrigrams := range langProfileList {
		if len(options.Whitelist) != 0 {
			//Skip non-whitelisted languages.
			if _, ok := options.Whitelist[lang]; !ok {
				continue
			}
		} else if len(options.Blacklist) != 0 {
			//skip blacklisted languages.
			if _, ok := options.Blacklist[lang]; ok {
				continue
			}
		}

		dist := calculateDistance(langTrigrams, trigrams)
		langDistances = append(langDistances, langDistance{lang, dist})
	}

	switch len(langDistances) {
	case 0:
		return -1, 0
	case 1:
		return langDistances[0].lang, 1
	default:
		return calculateConfidence(langDistances, trigrams)
	}
}

func calculateConfidence(langDistances []langDistance, trigrams map[string]int) (Lang, float64) {
	sort.SliceStable(langDistances, func(i, j int) bool { return langDistances[i].dist < langDistances[j].dist })
	langDist1 := langDistances[0]
	langDist2 := langDistances[1]
	score1 := maxTotalDistance - langDist1.dist
	score2 := maxTotalDistance - langDist2.dist

	var confidence float64
	if score1 == 0 {
		// If score1 is 0, score2 is 0 as well, because array is sorted.
		// Therefore there is no language to return.
		return -1, 0
	} else if score2 == 0 {
		// If score2 is 0, return first language, to prevent division by zero in the rate formula.
		// In this case confidence is calculated by another formula.
		// At this point there are two options:
		// * Text contains random characters that accidentally match trigrams of one of the languages
		// * Text really matches one of the languages.
		//
		// Number 500.0 is based on experiments and common sense expectations.
		confidence = float64((score1) / 500.0)
		if confidence > 1.0 {
			confidence = 1.0
		}
		return langDist1.lang, confidence
	}

	rate := float64((score1 - score2)) / float64(score2)

	// Hyperbola function. Everything that is above the function has confidence = 1.0
	// If rate is below, confidence is calculated proportionally.
	// Numbers 12.0 and 0.05 are obtained experimentally, so the function represents common sense.

	confidentRate := float64(12.0/float64(len(trigrams))) + 0.05
	if rate > confidentRate {
		confidence = 1.0
	} else {
		confidence = rate / confidentRate
	}

	return langDist1.lang, confidence
}

func calculateDistance(langTrigrams []string, textTrigrams map[string]int) int {
	var dist, totalDist int
	for i, trigram := range langTrigrams {
		if n, ok := textTrigrams[trigram]; ok {
			dist = abs(n - i)
		} else {
			dist = maxTrigramDistance
		}
		totalDist += dist
	}

	return totalDist
}

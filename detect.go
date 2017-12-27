package whatlanggo

import (
	"sort"
	"unicode"
)

const maxDist = 300

//Detect language and script of the given text.
func Detect(text string) Info {
	return DetectWithOptions(text, Options{})
}

//DetectLang detects only the language by a given text.
func DetectLang(text string) Lang {
	return Detect(text).Lang
}

//DetectLangWithOptions detects only the language of the given text with the provided options.
func DetectLangWithOptions(text string, options Options) Lang {
	return DetectWithOptions(text, options).Lang
}

//DetectWithOptions detects the language and script of the given text with the provided options.
func DetectWithOptions(text string, options Options) Info {
	script := DetectScript(text)
	if script != nil {
		lang := detectLangBaseOnScript(text, options, script)
		return Info{
			Lang:   lang,
			Script: script,
		}
	}
	return Info{}

}

func detectLangBaseOnScript(text string, options Options, script *unicode.RangeTable) Lang {
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
		return Cmn
	case unicode.Bengali:
		return Ben
	case unicode.Hangul:
		return Kor
	case unicode.Georgian:
		return Kat
	case unicode.Greek:
		return Ell
	case unicode.Kannada:
		return Kan
	case unicode.Tamil:
		return Tam
	case unicode.Thai:
		return Tha
	case unicode.Gujarati:
		return Guj
	case unicode.Gurmukhi:
		return Pan
	case unicode.Telugu:
		return Tel
	case unicode.Malayalam:
		return Mal
	case unicode.Oriya:
		return Ori
	case unicode.Myanmar:
		return Mya
	case unicode.Sinhala:
		return Sin
	case unicode.Khmer:
		return Khm
	case _HiraganaKatakana:
		return Jpn
	}
	return -1
}
func detectLangInProfiles(text string, options Options, langProfileList langProfileList) Lang {
	trigrams := getTrigramsWithPositions(text)
	type langDistance struct {
		lang Lang
		dist int
	}
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

	if len(langDistances) == 0 {
		return -1
	}
	sort.SliceStable(langDistances, func(i, j int) bool { return langDistances[i].dist < langDistances[j].dist })

	return langDistances[0].lang
}

func calculateDistance(langTrigrams []string, textTrigrams map[string]int) int {
	var dist, totalDist int
	for i, trigram := range langTrigrams {
		if n, ok := textTrigrams[trigram]; ok {
			dist = abs(n - i)
		} else {
			dist = maxDist
		}
		totalDist += dist
	}

	return totalDist
}

package whatlanggo

import (
	"sort"
	"strings"
	"unicode"
)

type trigram struct {
	trigram string
	count   int
}

//convert punctuations and digits to space.
func toTrigramChar(ch rune) rune {
	if isStopChar(ch) {
		return ' '
	}
	return ch
}

func getTrigramsWithPositions(text string) map[string]int {
	counterMap := count(text)
	trigrams := make([]trigram, len(counterMap))

	i := 0
	for tg, count := range counterMap {
		trigrams[i] = trigram{tg, count}
		i++
	}

	sort.SliceStable(trigrams, func(i, j int) bool {
		if trigrams[i].count == trigrams[j].count {
			return strings.Compare(trigrams[i].trigram, trigrams[j].trigram) < 0
		}
		return trigrams[i].count < trigrams[j].count
	})

	trigramsWithPositions := map[string]int{}

	j := 0
	for i := len(trigrams) - 1; i >= 0; i-- {
		trigramsWithPositions[trigrams[i].trigram] = j
		j++
	}
	return trigramsWithPositions
}

func count(text string) map[string]int {
	var r1, r2, r3 rune
	trigrams := map[string]int{}
	var txt []rune

	for _, r := range text {
		txt = append(txt, unicode.ToLower(toTrigramChar(r)))
	}
	txt = append(txt, ' ')

	r1 = ' '
	r2 = txt[0]
	for i := 1; i < len(txt); i++ {
		r3 = txt[i]
		if !(r2 == ' ' && (r1 == ' ' || r3 == ' ')) {
			trigram := []rune{}
			trigram = append(trigram, r1)
			trigram = append(trigram, r2)
			trigram = append(trigram, r3)
			if trigrams[string(trigram)] == 0 {
				trigrams[string(trigram)] = 1
			} else {
				trigrams[string(trigram)]++
			}
		}
		r1 = r2
		r2 = r3
	}

	return trigrams
}

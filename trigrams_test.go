package whatlanggo

import "testing"

func TestCount(t *testing.T) {
	tests := map[string]map[string]int{
		"":             {"": 0},
		",":            {"": 0},
		"a":            {" a ": 1},
		"-a-":          {" a ": 1},
		"yes":          {" ye": 1, "yes": 1, "es ": 1},
		"Give - IT...": {" gi": 1, "giv": 1, "ive": 1, "ve ": 1, " it": 1, "it ": 1},
	}

	for key, value := range tests {
		got := count(key)

		for key1, value1 := range value {
			if got[key1] != value1 {
				t.Fatalf("%s got %d want %d", key1, got[key1], value1)
			}
		}
	}
}

func TestToTrigramChar(t *testing.T) {
	tests := map[rune]rune{
		'a': 'a', 'z': 'z', 'A': 'A', 'Z': 'Z', 'Ж': 'Ж', 'ß': 'ß',
		//punctuation, digits, ... etc
		'\t': ' ', '\n': ' ', ' ': ' ', '.': ' ', '0': ' ', '9': ' ', ',': ' ', '@': ' ',
		'[': ' ', ']': ' ', '^': ' ', '\\': ' ', '`': ' ', '|': ' ', '{': ' ', '}': ' ', '~': ' '}

	for r, want := range tests {
		got := toTrigramChar(r)
		if got != want {
			t.Fatalf("%q got %q want %q", r, got, want)
		}
	}
}

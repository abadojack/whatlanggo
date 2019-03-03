package whatlanggo

import "testing"

func TestIsStopChar(t *testing.T) {
	tests := map[rune]bool{
		//Space
		'\t': true, '\n': true, '\v': true, '\r': true, '\f': true, 0x85: true, 0xA0: true,
		//Digits
		'0': true, '1': true, '2': true, '3': true, '5': true, '6': true, '9': true,
		//Punct
		';': true, '!': true, '"': true,
		//Symbol
		'`': true,
	}

	for r, want := range tests {
		got := isStopChar(r)
		if got != want {
			t.Fatalf("%v want %t got %t", r, want, got)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := map[int]int{
		1:      1,
		-0:     0,
		69:     69,
		-65535: 65535,
		65535:  65535,
	}

	for x, want := range tests {
		got := abs(x)
		if got != want {
			t.Fatalf("want %d got %d", want, got)
		}
	}
}

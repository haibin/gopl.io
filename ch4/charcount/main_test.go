package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			input: "hello world",
			counts: map[rune]int{
				'h': 1,
				'l': 3,
				'o': 2,
				'r': 1,
				'd': 1,
				'e': 1,
				' ': 1,
				'w': 1,
			},
			utflen:  [utf8.UTFMax + 1]int{0, 11},
			invalid: 0,
		},
	}

	for _, test := range tests {
		counts, utflen, invalid := charcount(strings.NewReader(test.input))
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("charcount(%s) returns Unicode chars %v, want %v", test.input, counts, test.counts)
		}
		if utflen != test.utflen {
			t.Errorf("charcount(%s) returns lengths %v, want %v", test.input, utflen, test.utflen)
		}
		if invalid != test.invalid {
			t.Errorf("charcount(%s) returns %d invalid, want %d", test.input, invalid, test.invalid)
		}
	}
}

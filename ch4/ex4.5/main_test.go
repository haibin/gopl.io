package ex4point5

import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	cases := []struct {
		input  []string
		output []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "a"}, []string{"a", "b", "a"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		r := dedupe(c.input)
		if !reflect.DeepEqual(r, c.output) {
			t.Errorf("dedupe(%v) == %v, want %v", c.input, r, c.output)
		}
	}
}

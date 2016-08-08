// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(132)
	x.Add(133)
	x.Add(239832)

	if x.Len() != 3 {
		t.Errorf("Len() return %d, want 3", x.Len())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(132)
	x.Add(133)
	x.Add(239832)
	if x.Len() != 3 {
		t.Errorf("Len() return %d, want 3", x.Len())
	}

	x.Remove(133)
	x.Remove(134)
	x.Remove(11111111)
	if x.Len() != 2 {
		t.Errorf("Len() return %d, want 2", x.Len())
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(132)
	x.Add(133)
	x.Add(239832)
	x.Clear()
	if x.Len() != 0 {
		t.Errorf("Len() return %d, want 0", x.Len())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(132)
	x.Add(133)
	x.Add(239832)

	s := x.Copy().String()
	want := "{132 133 239832}"
	if s != want {
		t.Errorf("Copy() returns %s, want %s", s, want)
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.Add(132)
	x.Add(133)
	x.Add(239832)

	x.AddAll(1, 3, 2)
	s := x.String()
	want := "{1 2 3 132 133 239832}"
	if s != want {
		t.Errorf("Copy() returns %s, want %s", s, want)
	}
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

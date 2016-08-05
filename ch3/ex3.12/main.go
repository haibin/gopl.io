package main

import (
	"fmt"
	"os"
)

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]

	fmt.Printf("%s %s %t\n", s1, s2, isAnagrams(s1, s2))
}

func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

    m1 := getMap(s1)
    m2 := getMap(s2)

    for k, v := range m1 {
        if m2[k] != v {
            return false
        }
    }
    
    return true
}

func getMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	return m
}

// Write a function that counts the number of bits that are different in two SHA256 hashes.
// (See PopCount from Section 2.6.2.)
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println("c1", popCount(c1))
	fmt.Println("c2", popCount(c2))
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popCount returns the population count (number of set bits) of x.
// sha256.Size is 32
func popCount(x [sha256.Size]byte) int {
	return int(pc[x[0]] + pc[x[1]] + pc[x[2]] + pc[x[3]])
}

// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]

    // TODO: Use command-line flag, not command-line argument.
	var flag string
	if len(os.Args) > 2 {
		flag = os.Args[2]
	}

	switch flag {
	case "384":
		fmt.Printf("SHA384 hash %x\n", sha512.Sum384([]byte(input)))
	case "512":
		fmt.Printf("SHA512 hash %x\n", sha512.Sum512([]byte(input)))
	default:
		fmt.Printf("SHA256 hash %x\n", sha256.Sum256([]byte(input)))
	}
}

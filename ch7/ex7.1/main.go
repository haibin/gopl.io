// Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter is
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
    *c += WordCounter(count)
	return count, nil
}

// LineCounter is
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// Default is ScanLines
	// Count the lines.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
    *c += LineCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // word count: 1

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // word count: 2

	var l LineCounter
    l.Write([]byte("hello"))
	fmt.Println(l) // word count: 1

	l = 0 // reset the counter
	var line2 = "Dolly\nDolly\nHaha"
	fmt.Fprintf(&l, "hello, %s", line2)
	fmt.Println(l) // word count: 2    
}

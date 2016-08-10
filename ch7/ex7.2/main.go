package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	counter int64
	w       io.Writer
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	c.counter += int64(len(p))
	return c.w.Write(p)
}

// CountingWriter returns
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	b := &ByteCounter{w: w}
	return b, &b.counter
}

func main() {
	byteCounter, count := CountingWriter(os.Stdout)
	fmt.Fprint(byteCounter, "hello")
    fmt.Println("\tcount", *count)
}

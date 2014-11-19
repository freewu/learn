package main

import (
	"fmt"
	"crypto/sha1"
	"io"
)

// echo sha1("123456"); // 7c4a8d09ca3762af61e59520943dc26494f8941b
func main() {
	// New
	h := sha1.New();
	io.WriteString(h,"123456")
	fmt.Printf("%x\r\n",h.Sum(nil))

	// Sum
	data := []byte("123456")
	fmt.Printf("%x\r\n",sha1.Sum(data))
}
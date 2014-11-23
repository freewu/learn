package main 

import (
	"fmt"
	"strings"
	"unicode"
)

// echo strpos("123456","23"); // 1
// echo strpos("123456","24"); // false
// echo strpos("123456","12"); // 0
func main() {
	// func Index(s, sep string) int
	// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
	fmt.Println(strings.Index("123456","23")) // 1
	fmt.Println(strings.Index("123456","24")) // -1
	fmt.Println(strings.Index("123456","12")) // 0

	// IndexAny 
	// func IndexAny(s, chars string) int
	// IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
	fmt.Println(strings.IndexAny("123456","2356")) // 1
	fmt.Println(strings.IndexAny("123456","78")) // 11
	fmt.Println(strings.IndexAny("123456","748")) // 3

	// func IndexByte(s string, c byte) int
	// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
	fmt.Println(strings.IndexByte("123456",'2')) // 1
	fmt.Println(strings.IndexByte("123456",'7')) // -1

	// func IndexFunc(s string, f func(rune) bool) int
	// IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.
	f := func(r rune) bool {
		return unicode.Is(unicode.Han,r)
	}
	fmt.Println(strings.IndexFunc("Hello,中国",f)) // 7
	fmt.Println(strings.IndexFunc("Hello,World",f)) // -1

	// func IndexRune(s string, r rune) int
	// IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s.
	fmt.Println(strings.IndexRune("123456",'2')) // 1
	fmt.Println(strings.IndexRune("123456",'7')) // -1

}
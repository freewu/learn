package main 

import (
	"fmt"
	"strings"
	"unicode"
)

// echo strrpos("1234561234","23"); // 7
// echo strrpos("1234561234","24"); // false
// echo strrpos("1234561234","12"); // 6

func main() {
	// func LastIndex(s, sep string) int
	// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
	fmt.Println(strings.LastIndex("1234561234","23")) // 7 
	fmt.Println(strings.LastIndex("1234561234","24")) // -1
	fmt.Println(strings.LastIndex("1234561234","12")) // 6

	// func LastIndexAny(s, chars string) int
	// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
	fmt.Println(strings.LastIndexAny("1234561234","24")) // 9

	// func LastIndexFunc(s string, f func(rune) bool) int
	// LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.	
	f := func(r rune) bool {
		return unicode.Is(unicode.Han,r)
	}
	fmt.Println(strings.LastIndexFunc("Hello,中国1111",f)) // 9
	fmt.Println(strings.LastIndexFunc("Hello,World",f)) // -1
}
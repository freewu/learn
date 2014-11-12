package main

import(
	"fmt"
	"strings"
	"unicode/utf8"
)

// echo strlen("123456"); // 6
// echo strlen("中国");   // 6
// echo mb_strlen("中国","utf8")； // 2
func main() {
	// 直接使用len
	fmt.Println(len("123456")) // 6
	fmt.Println(len("中国")) // 6 
	// 使用strings
	fmt.Println(strings.Count("123456","")) // 7  算上了\0
	fmt.Println(strings.Count("中国","")) // 3 算上了\0  能识别utf8
	// 使用unicode/utf8
	// http://stackoverflow.com/questions/12668681/go-language-string-length
	fmt.Println(utf8.RuneCountInString("123456")) // 6
	fmt.Println(utf8.RuneCountInString("中国")) // 2
}
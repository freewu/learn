package main

import(
	"fmt"
	"strings"
)

// echo trim("  12 3 4   "); // "12 3 4"
// echo ltrim("  12 3 4   "); // "12 3 4   "
// echo rtrim("  12 3 4   "); // "  12 3 4"
// echo trim("  中 国   "); // "中 国"
// echo ltrim("  中 国   "); // "中 国   "
// echo rtrim("  中 国  "); // "  中 国"

// echo trim("!!123!4!!!!","!"); // "123!4"

// echo trim("!!!12x3xxx","!x");
func main() {
	fmt.Println(strings.Trim("  12 3 4   "," ")) // 
	fmt.Println(strings.TrimLeft("  12 3 4   "," ")) // 
	fmt.Println(strings.TrimRight("  12 3 4   "," ")) // 
	fmt.Println(strings.Trim("  中 国   "," ")) // 
	fmt.Println(strings.TrimLeft("  中 国   "," ")) // 
	fmt.Println(strings.TrimRight("  中 国   "," ")) // 

	fmt.Println(strings.Trim("!!123!4!!!!","!")) // 

	fmt.Println(strings.Trim("!!!12x3xxx","!x")) // 
}
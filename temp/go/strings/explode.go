package main

import(
	"fmt"
	"strings"
)

// var_dump(explode(",","1,2,3"));
func main() {
	fmt.Printf("%q\r\n",strings.Split("1,2,3",","))

	// keep the slices string
	fmt.Printf("%q\r\n",strings.SplitAfter("1,2,3",","))
}
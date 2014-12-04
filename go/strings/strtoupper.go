package main

import (
	"fmt"
	"strings"
)

// echo strtoupper("aBcD"); // ABCD
func main() {
	fmt.Println(strings.ToUpper("aBcD"))

	// ToTitle 可以处理unicode的大写 （欧洲的其它国家的大小）
	fmt.Println(strings.ToTitle("aBcD"))
	fmt.Println(strings.ToTitle("хлеб"))
}
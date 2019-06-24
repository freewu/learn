package main

import (
	"fmt"
	"strings"
)

// str_replace("x","0","1234xxx345");
func main() {
	fmt.Println(strings.Replace("1234xxx345","x","0",-1))
	// 只替换开头2个x
	fmt.Println(strings.Replace("x1234xxx345","x","0",2))

	// 使用Replacer
	r := strings.NewReplacer("<","&lt;",">","&gt;")
	fmt.Println(r.Replace("<html>"))
}
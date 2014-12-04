package main

import (
	"fmt"
	"strings"
)

// todo: this code is ugly need refactor when learn more 
func Ucfirst(str string) string {
	rs := []rune(str)
	rl := len(rs)

	f := string(rs[0:1])
	f = strings.ToUpper(f)
	//fmt.Println(strings.ToUpper(f))
	rf := []rune(f)
	rs[0] = rf[0]
	// rs[0] = strings.ToUpper(f)
	return string(rs[0:rl])
}

// echo ucfirst("hello world"); // Hello world
func main() {
	fmt.Println(Ucfirst("hello world"))
}
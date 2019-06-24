package main

import (
	"fmt"
	//"strings"
)

func Substr(str string,start,length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start,end = end,start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// substr("1234",0,2); // 12
// go 居然没有这个方法
func main() {
	fmt.Println(Substr("1234",0,2))
}
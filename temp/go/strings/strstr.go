package main

import (
	"fmt"
	"strings"
)

// echo strstr("www@163.com","@"); // @163.com
func main() {
	s := strings.SplitAfterN("www@163.com","@",2);
	//fmt.Printf("%q\r\n",s)

	// echo strstr("www@163.com","@"); // @163.com
	fmt.Println(s[1])

	// echo strstr("www@163.com","@",true); // www
	fmt.Println(strings.Trim(s[0],"@"))
}
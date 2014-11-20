package main

import (
	"fmt"
	"net/url"
)

// echo urlencode("中国"); // %E4%B8%AD%E5%9B%BD
// echo urldecode("%E4%B8%AD%E5%9B%BD"); // 中国
func main() {
	// encode
	e := url.QueryEscape("中国")
	fmt.Println(e); // %E4%B8%AD%E5%9B%BD

	// decode
	str,err := url.QueryUnescape("%E4%B8%AD%E5%9B%BD")
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	fmt.Println(str)
}
package main

import(
	"fmt"
	"crypto/md5"
	"io"
)

// echo md5("1234"); // 81dc9bdb52d04dc20036dbd8313ed055
func main() {
	h := md5.New()
	io.WriteString(h,"1234")
	fmt.Printf("%x\r\n",h.Sum(nil))
	//fmt.Println(string(h.Sum(nil)))
	// byte
	data := []byte("1234")
	fmt.Printf("%x\r\n",md5.Sum(data))
	//fmt.Println(string(md5.Sum(data)))
}
package main

import (
	"fmt"
	"encoding/base64"
)

// echo base64_encode("中国"); // 5Lit5Zu9
// echo base64_decode("5Lit5Zu9"); // 中国
func main() {
	// encode
	data := []byte("中国")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str); // 5Lit5Zu9

	// decode
	data,err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	fmt.Printf("%q\n",data)
	fmt.Println(string(data))
}
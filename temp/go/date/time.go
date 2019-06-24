package main

import (
	"fmt"
	"time"
)

// echo time();
func main() {
	fmt.Println(time.Now().Unix())
}
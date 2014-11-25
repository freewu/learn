package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	if rand.Intn(100) > 50 {
		fmt.Println("great than 50")
	} else {
		fmt.Println("less than 50")
	}
}
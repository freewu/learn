package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan <- string,msg string) {
	pings <- msg
}

//  <- chan string  & chan <- string 晕了
func pong(pings <- chan string,pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	// == channel
	message := make(chan string)

	go func() { message <- "ping"}()

	msg := <-message
	fmt.Println(msg)

	// == channel buffer
	msg1 := make(chan string,2)
	msg1 <- "buffered"
	msg1 <- "channel"
	//msg1 <- "channel1"

	fmt.Println(<- msg1)
	fmt.Println(<- msg1)
	//fmt.Println(<- msg1)

	// == channel synchronization
	done := make(chan bool,1)
	go worker(done)

	// 被注释了就没有运行了
	<- done

	// == channel directions
	pings := make(chan string,1)
	pongs := make(chan string,1)

	ping(pings,"passed msg")
	pong(pings,pongs)
	fmt.Println(<-pongs)

}
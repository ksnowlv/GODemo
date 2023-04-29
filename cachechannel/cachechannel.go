package main

import (
	"fmt"
	"time"
)

func main() {
	channelTest()
}

const LOOPNUM = 100

func writeChannel(ch chan int) {

	for i := 0; i < LOOPNUM; i++ {
		ch <- i

		if i%3 == 0 {
			time.Sleep(2 * time.Second)
			fmt.Println("sleep 2s")
		}
	}

	close(ch)
}

func channelTest() {

	ch := make(chan int, 2)

	go writeChannel(ch)

	for v := range ch {
		fmt.Println("receive:", v)
	}
}

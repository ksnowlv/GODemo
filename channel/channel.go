package main

import "fmt"

func main() {
	twoChannelTest()
	singleChannelTest()
	channelWithCloseTest()
}

// 双向信号
func helloChannel(channel chan int) {
	// v := <-channel
	// fmt.Printf("helloChannel read %d", v)
	fmt.Printf("helloChannel write %d\n", 10)
	channel <- 10
}

// 双向通道测试
func twoChannelTest() {
	// var num chan int
	var num chan int
	if num == nil {

		fmt.Println("Hello channel!")

		num = make(chan int)
		fmt.Printf("Type of num is %T\n", num)

	}

	go helloChannel(num)

	res := <-num

	fmt.Println("main rev num:", res)
}

// 单向通道
func singleChannel(send chan<- int, value int) {
	send <- value
}

// 单向通道测试
func singleChannelTest() {
	fmt.Println("---singleChannelTest")
	send := make(chan int)

	go singleChannel(send, 2)

	res := <-send
	fmt.Printf("read value:%d\n", res)
}

// 通道关闭测试
const LOOPNUM int = 100

func channelWithClose(channel chan int) {

	for i := 0; i < LOOPNUM; i++ {
		channel <- i
	}
	//如不关闭，主协程中，会死循环。
	close(channel)
}

func channelWithCloseTest() {
	fmt.Println("--channelWithCloseTest")
	ch := make(chan int)

	go channelWithClose(ch)

	for v := range ch {
		fmt.Println("receive:", v)
	}
}

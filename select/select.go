package main

import (
	"fmt"
	"time"
)

func main() {
	SelectTest()
}

func ServerInBeiJing(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Server In Beijing"
}

func ServerInShangHai(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Server In ShangHai"
}

func SelectTest() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ServerInBeiJing(ch1)
	go ServerInShangHai(ch2)

	//添加上该句，随机选择；不加此句，选择res2
	time.Sleep(3 * time.Second)

	select {
	case res1 := <-ch1:
		fmt.Println(res1)
	case res2 := <-ch2:
		fmt.Println(res2)
	}

	fmt.Println("main end")
}

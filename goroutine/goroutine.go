package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func sayHello() {
	gid := GetGid()
	fmt.Printf("child goruntine1 gid:%v \n", gid)
	fmt.Println("Hello goroutine\n")
}

func sayGoodbye() {
	gid := GetGid()
	fmt.Printf("child goruntine1 gid:%v \n", gid)
	fmt.Println("goodbye goroutine\n")
}

func main() {

	go sayHello()
	go sayGoodbye()
	time.Sleep(1 * time.Second)
	gid := GetGid()
	fmt.Printf("main goruntine1 gid:%v \n", gid)
}

func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

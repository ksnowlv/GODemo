package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitgroup()
}

func waitgroup() {

	num := 10

	var wg sync.WaitGroup

	//直接设定需要完成的工作任务数，不能大于实际工作任务数，不然，直接死锁。例如设定num+1，直接死锁
	wg.Add(num)
	for i := 0; i < num; i++ {
		//计数增1
		//wg.Add(1)
		go process(i, &wg)
	}

	//等待WaitGroup任务处理，当计数为0时，处理完成，否则阻塞等待处理。
	defer wg.Wait()
	fmt.Println("All routines done!!!")
}

func process(num int, wg *sync.WaitGroup) {
	fmt.Println("start Goroutine:", num)
	time.Sleep(time.Second)
	fmt.Printf("Goroutine end: %d\n", num)
	//计数减1
	wg.Done()
}

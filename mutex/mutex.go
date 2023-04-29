package main

import (
	"fmt"
	"sync"
)

var g_count = 0
var mutex sync.Mutex

func incrementWithMutex(wg *sync.WaitGroup) {
	mutex.Lock()
	g_count += 1
	mutex.Unlock()
	wg.Done()
}

const LOOPNUM = 1000

func mutexTest() {
	var wg sync.WaitGroup
	wg.Add(LOOPNUM)
	for i := 0; i < LOOPNUM; i++ {
		go incrementWithMutex(&wg)
	}
	wg.Wait()
	fmt.Println("final value of g_count", g_count)
}

func incrementWithChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	g_count += 1
	<-ch
	wg.Done()
}

func mutexTestForChannel() {

	//用通道实现互斥
	ch := make(chan bool, 1)

	var wg sync.WaitGroup
	wg.Add(LOOPNUM)

	for i := 0; i < LOOPNUM; i++ {
		go incrementWithChannel(&wg, ch)
	}
	wg.Wait()
	fmt.Println("final value of g_count", g_count)
}

//特殊说明！！！
//就上述场景而言，使用Mutex实现互斥更佳，不需要协程间的通信。
func main() {
	mutexTest()
	mutexTestForChannel()
}

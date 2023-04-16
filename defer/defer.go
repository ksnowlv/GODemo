package main

import "fmt"

// defer调用时，GO会把defer调用放入到栈中，按照LIFO（后进先出）的顺序执行

func test1() {
	fmt.Println("test1")
}

func test2() {
	fmt.Println("test2")
}

func test3() {
	fmt.Println("test3")
}

func main() {

	fmt.Println("---test defer---")
	defer test1()
	fmt.Println("test---x1")
	defer test2()
	fmt.Println("test---x2")
	defer test3()
	fmt.Println("test---x3")
}

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello world!!!")

	var age int = 10
	var name = "ksnowlv"

	fmt.Println("age:", age, ",name:", name)

	var (
		address = "beijing"
		width   = 1.2
	)
	fmt.Println("address:", address, ",width:", width)

	num, flag := 1, 2

	fmt.Println("num:", num, ",flag:", flag)

	a, b := 10.2, 8.1
	c := math.Min(a, b)
	d := math.Max(a, b)

	fmt.Println("a,b中小的数:", c, "a,b中大的数:", d)
}

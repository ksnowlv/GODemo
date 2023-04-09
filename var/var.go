package main

import (
	"fmt"
	"math"
)

func main() {

	var age int = 10
	var name = "ksnowlv"
	fmt.Println("age:", age, ",name:", name)

	age = 20
	fmt.Println("age:", age, ",name:", name)

	var (
		address = "beijing"
		width   = 1.2
		height  int
	)
	fmt.Println("address:", address, ",width:", width, ",height", height)

	var num, flag int = 1, 2
	fmt.Println("num:", num, ",flag:", flag)

	a, b := 10.2, 8.1
	c := math.Min(a, b)
	d := math.Max(a, b)
	fmt.Println("a,b中小的数:", c, "a,b中大的数:", d)
}

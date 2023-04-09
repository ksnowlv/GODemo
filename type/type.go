package main

import "fmt"

func main() {

	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b

	fmt.Println("d:", d)

	e, f := 5.113, 9.218
	fmt.Println("type of e %T f %T\n", e, f)
	sum := e + f
	dif := e - f
	fmt.Printf("e + f:%.2f,e - f:%.2f\n", sum, dif)

	g := 10
	j := float64(g)
	fmt.Println("j:", j)

	c1 := complex(1, 3)
	c2 := 1 + 2i
	res1 := c1 + c2
	res2 := c1 - c2

	fmt.Println("c1 + c2:", res1, "c1 - c2:", res2)

	firstName := "wei"
	lastName := "lv"

	fullName := firstName + lastName

	fmt.Println("fullName:", fullName)

}

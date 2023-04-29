package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("---test for if")
	total := 120
	res := totalResult(total)
	fmt.Println(res)

	total = 80
	res = totalResult(total)
	fmt.Println(res)

	total = 30
	res = totalResult(total)
	fmt.Println(res)

	fmt.Println("---test for for")
	fmt.Println(forTest())

	fmt.Println("---test for switch")
	fmt.Println(swichTest())
}

func totalResult(total int) string {
	if total >= 100 {
		return "total >= 100"
	} else if total >= 50 {
		return "total >= 50"
	} else {
		return "total < 50"
	}
}

func forTest() string {
	const loopNum = 10

	res := ""

	for i := 0; i < loopNum; i++ {
		res += strconv.Itoa(i) + " "
	}

	res += "\n"

	for i := 0; i < loopNum; i++ {
		res += strconv.Itoa(i) + " "

		if i >= 5 {
			break
		}
	}

	res += "\n"

	for i := 0; i < loopNum; i++ {
		if i == 5 || i == 7 {
			continue
		}
		res += strconv.Itoa(i) + " "
	}

	return res
}

func swichTest() string {
	res := ""

	total := 2

	switch total {
	case 0:
		res += "0" + " "
	case 1:
		res += "1" + " "
	case 2:
		res += "2" + " "
	case 3:
		res += "2" + " "
	default:
		res += "default: "
	}

	res += "\n"

	letter := "c"

	switch letter {
	case "a":
		res += "a "
	case "b", "c", "d", "e":
		res += "命中bcde了 "
	default:
		res += "letter default "
	}

	res += "\n"

	num := 1

	switch {
	case num >= 50:
		res += "num >=50 "
		fallthrough

	case num >= 10:
		res += "num >=10 "
		fallthrough

	case num >= 5:
		res += "num >=5 "
		fallthrough

	case num < 5:
		res += "num < 5 "
	}

	return res
}

package main

import "fmt"

type Student struct {
	firstName   string
	lastName    string
	age         int
	sex         string
	phoneNumber string
}

type Address struct {
	province string
	city     string
}

type Person struct {
	name    string
	address Address
}

func main() {

	student1 := Student{
		firstName:   "wei",
		lastName:    "lv",
		age:         30,
		sex:         "male",
		phoneNumber: "15100012345",
	}

	fmt.Println("student 1:", student1)

	student2 := Student{
		"wei",
		"lv",
		30,
		"male",
		"15100012345",
	}

	fmt.Println("student 2:", student2)

	var student3 Student
	fmt.Println("student 3:", student3)

	student4 := &student1
	fmt.Println("student 4:", *student4)

	p := Person{
		name: "ksnow",
		address: Address{
			"Beijing",
			"Beijing",
		},
	}

	fmt.Println("Person :", p)
}

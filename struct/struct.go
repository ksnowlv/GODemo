package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	firstName   string
	lastName    string
	age         int
	sex         string
	phoneNumber string
}

// 类的方法
func (s Student) showStudentInfo() {

	res := "firstName:" + s.firstName +
		",lastName:" + s.lastName +
		",age:" + strconv.Itoa(s.age) +
		",sex:" + s.sex +
		",phoneNumber:" + s.phoneNumber
	fmt.Println(res)
}

// 函数
func showStudentInfo(s Student) {
	res := "firstName:" + s.firstName +
		",lastName:" + s.lastName +
		",age:" + strconv.Itoa(s.age) +
		",sex:" + s.sex +
		",phoneNumber:" + s.phoneNumber
	fmt.Println(res)
}

// 使用值接收器的方法
func (s Student) setAge(age int) {
	s.age = age
}

// 使用指针接收器的方法，当拷贝一个结构体对象代价过于昂贵时
func (s *Student) setFirstName(name string) {
	s.firstName = name
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
	student1.showStudentInfo()
	showStudentInfo(student1)

	student2 := Student{
		"wei",
		"lv",
		30,
		"male",
		"15100012345",
	}

	fmt.Println("student 2:", student2)
	student2.setAge(20)
	student2.setFirstName("weiwei")
	student2.showStudentInfo()

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

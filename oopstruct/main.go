package main

import "fmt"

type Student struct {
	name  string
	phone string
	age   int
}

func New(name string, phone string, age int) Student {
	return Student{name, phone, age}
}

func (s *Student) ShowInfo() {
	fmt.Printf("name:%s,phone:%s,age:%d\n",
		s.name,
		s.phone,
		s.age)
}

func main() {
	s := Student{
		"lvwei",
		"152",
		30,
	}
	s.ShowInfo()

	s = New("ksnowlv", "151", 31)
	s.ShowInfo()
}

package student

import "fmt"

type Student struct {
	name  string
	phone string
	age   int
}

func (s *Student) ShowInfo() {
	fmt.Printf("name:%s,phone:%s,age:%d\n",
		s.name,
		s.phone,
		s.age)
}

package main

import (
	. "fmt"
	"testperson/person"
)

//开启go

func main() {

	Println("person test")
	person.ShowLog()

	p := person.Person{Name: "ksnowlv", Phone: "152", Age: 30}

	p.ShowInfo()
}

package main

import (
	. "fmt"
	"testperson/person"
)

//开启go

func main() {

	Println("person test")
	person.ShowLog()

	p := person.Person{"ksnowlv", "152", 30}

	p.ShowInfo()
}

package person

import "fmt"

// 在golang中，软件结构是以包为单位的，在同一个包内属于内部，不同包之间属于包间。
// 给外部包用的变量，必须首字母大写，否则就会出现上述问题。
// 同时，结构体内部的变量，如果不声明为首字母大写的变量，也会出现该问题。
type Person struct {
	Name  string
	Phone string
	Age   int
}

func ShowLog() {
	fmt.Println("showlog")
}

func (p Person) ShowInfo() {
	fmt.Printf("name:%s, phone:%s, age:%d\n",
		p.Name,
		p.Phone,
		p.Age)
}

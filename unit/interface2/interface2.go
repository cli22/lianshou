package interface2

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintName() {
	fmt.Println(s.Name)
}

type SameStudent interface {
	PrintName()
}



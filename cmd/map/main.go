package main

import (
	//. "fmt"
	//"lianshou/unit/map2"
	"lianshou/unit/interface2"
)

//func panicfunc() {
//	panic("just test")
//}

func main() {
	//map2.InitMap()
	//map2.MakeMap()
	//map2.Defer()
	//map2.Filter()
	//map2.ThrowPanic(panicfunc)
	var ss interface2.SameStudent

	s1 := interface2.Student{"111", 111}

	ss = s1
	ss.PrintName()
	//s := interface2.Student{"lcy", 22}
	//s.PrintName()
}

package map2

import (
	"fmt"
)

func InitMap() {
	rating := map[string]float32{"C": 5, "Go": 4, "C++": 3}

	if csharp, ok := rating["C#"]; ok {
		fmt.Println(csharp)
	} else {
		fmt.Println("404")
	}
}

func MakeMap() {
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m)
	//mn := new(map[string]int)
	i := new(int)
	i1 := *i
	i1 = 5
	fmt.Println(i1)
}

func Defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

type testInt func(int) bool

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func filter(f testInt, i int) bool {
	if f(i) {
		return true
	}

	return false
}

func Filter() {
	for i := 0; i < 5; i++ {
		fmt.Println(filter(isOdd, i))
	}
}

var user = ""

func init() {
	if user == "" {
		panic("no user")
	}
	fmt.Println("init")
}

func ThrowPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

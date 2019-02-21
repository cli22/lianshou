package main

import "fmt"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 2
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main()  {
	fmt.Println(f1())
	fmt.Println(f2())
}
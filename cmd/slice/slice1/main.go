package main

import "fmt"

func main() {
	var  s []int
	s = []int{1,2,3}
	//for i := 1; i <= 3; i++ {
	//	s = append(s, i)
	//}
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

package main

import "fmt"

//func main() {
//	var s []int
//	for i := 1; i <= 3; i++ {
//		s = append(s, i)
//	}
//	reverse(s)
//	fmt.Println("after", len(s))
//	fmt.Println(s[:4])
//}
//
//func reverse(s []int) {
//	s = append(s, 999)
//	fmt.Println(len(s))
//	for i, j := 0, len(s)-1; i < j; i++ {
//		j = len(s) - (i + 1)
//		s[i], s[j] = s[j], s[i]
//	}
//	fmt.Println("befor", s)
//}
//
//func main() {
//	a := []int{7, 8, 9}
//	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
//	ap(a)
//	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
//}
//
//func ap(a []int) {
//	a = append(a, 10)
//}

func main() {
	a := []int{7,8,9}
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	ap(a)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
}

func ap(a []int) {
	a[0] = 1
	a = append(a, 10)
}


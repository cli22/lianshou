package main

import (
	"lianshou/cmd/offer/common"
	"fmt"
)

func printReverse(head *common.ListNode) {
	if head == nil {
		return
	}
	if head.Next != nil {
		printReverse(head.Next)
	}
	fmt.Println(head.Val)
}

func main() {
	//head1 := &common.ListNode{Val: 1, Next: &common.ListNode{2, nil}}
	//head1 := &common.ListNode{Val: 1, Next: nil}
	head1 := &common.ListNode{Val: 1, Next: &common.ListNode{2, &common.ListNode{3, nil}}}
	//printReverse(nil)
	printReverse(head1)
}

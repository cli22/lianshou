package main

import (
	"fmt"

	"lianshou/cmd/offer/common"
)

func AddToTail(head **common.ListNode, value int) {
	node := &common.ListNode{Val: value}
	if *head == nil {
		*head = node
	} else {
		phead := *head
		for {
			if phead.Next != nil {
				phead = phead.Next
			} else {
				phead.Next = node
				break
			}

		}

	}
	fmt.Println(head)
}

func main() {
	//head1 := &common.ListNode{Val: 1, Next: &common.ListNode{2, nil}}
	//AddToTail(&head1, 0)

	var head2 *common.ListNode
	AddToTail(&head2, 0)

	fmt.Println(head2)
}

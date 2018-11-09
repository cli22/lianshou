package main

import (
	"lianshou/cmd/offer/common"
	"fmt"
)

func remove(head **common.ListNode, value int) {
	if head == nil || *head == nil {
		return
	}
	var toDeleted *common.ListNode = nil
	if (*head).Val == value {
		toDeleted = *head
		*head = (*head).Next
	} else {
		phead := *head
		for {
			if phead.Next != nil && phead.Next.Val != value {
				phead = phead.Next
			} else if phead.Next != nil && phead.Next.Val == value {
				toDeleted = phead.Next
				phead.Next = phead.Next.Next

			} else {
				break
			}
		}
	}

	if toDeleted != nil {
		toDeleted = nil
	}
}

func main() {
	head1 := &common.ListNode{Val: 1, Next: &common.ListNode{2, nil}}
	remove(&head1, 2)
	fmt.Println(head1)
}

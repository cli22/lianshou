package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	var carry int
	for {
		var x, y int
		if l1 != nil || l2 != nil {
			if l1 == nil {
				x = 0
			} else {
				x = l1.Val
			}

			if l2 == nil {
				y = 0
			} else {
				y = l2.Val
			}

			sum := x + y + carry
			if sum >= 10 {
				carry = 1
				sum = sum % 10
			} else {
				carry = 0
			}
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		} else {
			if carry > 0 {
				cur.Next = &ListNode{Val: carry}
			}
			break

		}

	}

	return res.Next
}

func main() {
	//l1 := &ListNode{2, &ListNode{6, &ListNode{3, nil}}}
	//l1 := &ListNode{2, &ListNode{3, nil}}
	l1 := &ListNode{Val: 5,}
	//l2 := &ListNode{4, &ListNode{5, &ListNode{6, nil}}}
	//l2 := &ListNode{4, &ListNode{5, nil}}
	l2 := &ListNode{Val: 5,}
	fmt.Println(addTwoNumbers(l1, l2))
}

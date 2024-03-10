package utils

import (
	"awesomeProject/pkg/structure"
	"fmt"
)

func MiddleNode(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func ListFromVals(vals []int) *structure.ListNode {
	if len(vals) == 0 {
		return nil
	}

	newHead := &structure.ListNode{}
	cur := newHead
	for _, val := range vals {
		cur.Next = &structure.ListNode{Val: val}
		cur = cur.Next
	}

	return newHead.Next
}

func PrintList(head *structure.ListNode) {
	if head == nil {
		return
	}
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

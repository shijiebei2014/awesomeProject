package heaps

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := ListNode{Val: vals[0]}
	cur := &head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{
			Val: vals[i],
		}
		cur = cur.Next
	}

	return &head
}

func print(head *ListNode) {
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

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	}
	l2.Next = mergeList(l2.Next, l1)
	return l2
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := middleNode(head)
	var (
		prev  *ListNode
		after *ListNode
	)
	prevHead, afterHead, cur := prev, after, head
	for cur != nil {
		if middle == cur {
			cur = cur.Next
			continue
		} else if cur.Val < middle.Val {
			if prev == nil {
				prevHead = &ListNode{Val: cur.Val}
				prev = prevHead
			} else {
				prev.Next = &ListNode{Val: cur.Val}
				prev = prev.Next
			}
		} else {
			if after == nil {
				afterHead = &ListNode{Val: cur.Val}
				after = afterHead
			} else {
				after.Next = &ListNode{Val: cur.Val}
				after = after.Next
			}
		}
		cur = cur.Next
	}

	print(prevHead)
	print(afterHead)
	newPrev := sortList(prevHead)
	newAfter := sortList(afterHead)
	return mergeList(mergeList(newPrev, &ListNode{Val: middle.Val}), newAfter)
}

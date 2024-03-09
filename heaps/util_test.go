package heaps

import "testing"

func Test_reverseList(t *testing.T) {
	head := toList([]int{1, 2, 3})
	newHead := reverseList(head)
	print(newHead)
}

func Test_reverseList2(t *testing.T) {
	head := toList([]int{1, 2, 3})
	newHead := &ListNode{Next: head}
	reverseList2(newHead, head)
	print(newHead.Next)
}

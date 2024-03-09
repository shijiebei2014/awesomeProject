package heaps

import "fmt"

/*
	1 -> 2 -> 3 -> 4

=>  1 -> 4 -> 2 -> 3
*/
func recordList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 1. 获取中间节点,2
	middle := middleNode2(head)
	// 2. 将2后面的反转
	reverseList2(middle, middle.Next)
	//print(head)
	// 3. 插入
	left := head
	right := middle.Next
	//fmt.Printf("right:%d\n", right.Val)
	// 1 2 3 6 5 4
	// 1 6 2 3 5 4
	for right != nil {
		next := right.Next
		middle.Next = right.Next
		right.Next = left.Next
		left.Next = right

		left = right.Next
		right = next
	}

	return head
}

func recordList22(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 1. 中间节点
	middle := middleNode22(head)
	fmt.Printf("%d\n", middle.Val)
	// 2. 反转
	reverseList22(middle, middle.Next)
	// 3. 插入
	left, right := head, middle.Next
	// 1 2 4 3
	// 1 4 2 3
	for right != nil {
		next := right.Next
		middle.Next = right.Next
		right.Next = left.Next
		left.Next = right

		left = left.Next
		right = next
	}
	return head
}

func middleNode22(head *ListNode) *ListNode {
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

func reverseList22(prev, head *ListNode) {
	// 2 -> 3
	// 3 -> 2
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return
}

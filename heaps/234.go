package heaps

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	//fmt.Printf("1\n")
	// 1. 找中间节点
	middle := middleNode33(head)
	//fmt.Printf("2 middle: %d %v\n", middle.Val, middle.Next)

	// 2. 反转
	reverseList33(middle, middle.Next)
	//print(middle)

	head1, head2 := head, middle.Next
	for head1 != middle {
		if head1.Val == head2.Val {
			head1 = head1.Next
			head2 = head2.Next
		} else {
			return false
		}
	}
	if head1 == middle {
		if head2 != nil && head1.Val != head2.Val {
			return false
		}
	}
	return true
}

func middleNode33(head *ListNode) *ListNode {
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

func reverseList33(prev, head *ListNode) {
	//fmt.Printf("%d %d\n", prev.Val, head.Val)
	// 2 -> 1 -> 3
	// 3 -> 1 -> 2
	right := head

	for right.Next != nil {
		next := right.Next
		//fmt.Printf("%d %v\n", right.Val, next)
		right.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
}

package heaps

/*
1 -> 2 -> 3
3 -> 2 -> 1
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Next: head}

	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = newHead.Next
		newHead.Next = next
	}
	return newHead.Next
}

func reverseList2(prev *ListNode, head *ListNode) {
	if head == nil {
		return
	}

	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	print(prev)
}

func toList(vals []int) *ListNode {
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

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func middleNode2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

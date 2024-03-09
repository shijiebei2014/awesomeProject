package heaps

import "testing"

func TestMiddleNode(t *testing.T) {
	head := generateList([]int{1, 2, 3, 4, 5})
	//print(head)
	node := middleNode(head)
	t.Log(node.Val)
}

func TestMergeLists(t *testing.T) {
	l1 := generateList([]int{2, 1})
	l2 := generateList([]int{5, 4, 3})
	node := mergeList(l1, l2)
	print(node)
}

func TestSortLists(t *testing.T) {
	l1 := generateList([]int{2, 1})
	l2 := generateList([]int{5, 4, 3})
	node := sortList(mergeList(l1, l2))
	print(node)
}

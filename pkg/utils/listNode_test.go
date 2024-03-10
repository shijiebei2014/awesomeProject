package utils

import (
	"testing"
)

func TestListFromVals(t *testing.T) {
	head := ListFromVals([]int{1, 2, 3})
	PrintList(head)
	t.Log(MiddleNode(head).Val)
}

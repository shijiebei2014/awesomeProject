package utils

import (
	"testing"
)

func TestToTree(t *testing.T) {
	node := ToTree([]int{1, 0, 2, 3, 4})
	//t.Log(node)
	PrintTree(node)
}

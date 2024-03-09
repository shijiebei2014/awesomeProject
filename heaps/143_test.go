package heaps

import (
	"testing"
)

func Test_recordList(t *testing.T) {
	head := toList([]int{1, 2, 3, 4})

	print(recordList(head))
	// 1 2 3 4 5 6
	// 1 2 3 6 5 4
	// 1 6 2 5 3 4
	head = toList([]int{1, 2, 3, 4, 5, 6})

	print(recordList(head))
}

func Test_recordList22(t *testing.T) {
	head := toList([]int{1, 2, 3, 4})
	print(recordList22(head))
	// 1 2 3 4 5 6
	// 1 2 3 6 5 4
	// 1 6 2 5 3 4
	head = toList([]int{1, 2, 3, 4, 5, 6})
	print(recordList(head))
}

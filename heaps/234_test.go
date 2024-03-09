package heaps

import "testing"

func Test_isPalindrome(t *testing.T) {
	head := toList([]int{3, 1, 2, 2, 1, 3})
	t.Log(isPalindrome(head))

	head = toList([]int{1, 2, 3, 4})
	t.Log(isPalindrome(head))
}

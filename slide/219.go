package slide

func containNearByDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || k == 0 {
		return false
	}

	left, right := 0, 0

	for left < len(nums) {
		if right+1 < len(nums) {
			if nums[left] == nums[right+1] {
				return true
			} else if right+1-left < k {
				right++
			} else {
				left++
				right = left
			}
		} else {
			left++
			right = left
		}
	}

	return false
}

package slide

import "fmt"

func longestMountain(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	left, right, res := 0, 1, 0
	asc := arr[left] < arr[right]

	for left < len(arr)-1 {
		// 递增
		if arr[left] < arr[left+1] {
			if !asc { // 递减到递增
				asc = true
				right = left + 1
			} else { // 递增
				right++
			}
		} else if arr[left] == arr[left+1] {
			left++
			right = left
		} else {     // 递减
			if asc { // 从递增到递减
				asc = false
				right++
			} else {
				right++
			}
		}
	}

	return res
}

// 增减
func longestMountain2(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	var diffs = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			diffs[0] = 0
		} else {
			diffs[i] = arr[i] - arr[i-1]
		}
	}
	fmt.Println(diffs)
	left, right := 0, len(diffs)-1
	for left < right {
		// 找到第一个开始递增
		for left <= right && diffs[left] < 0 {
			left++
		}
		// 找到最后一个个开始递减
		for right >= left && diffs[right] > 0 {
			right--
		}

		l, r := left, right
		for l+1 < len(diffs) && diffs[l+1] > 0 {
			l++
		}
		for r-1 >= 0 && diffs[r-1] < 0 {
			r--
		}
		if r > 0 && arr[r] < arr[r-1] {
			r--
		}
		fmt.Printf("left:%d right:%d l:%d r:%d\n", left, right, l, r)
		if r-l <= 1 && r >= l {
			if left == 0 {
				return right - left + 1
			}
			return right - left + 2
			break
		}

		left++
	}

	return 0
}

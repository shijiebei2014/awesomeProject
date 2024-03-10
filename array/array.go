package array

func mergeSorted(arr1, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	arr1 = append(arr1, make([]int, n)...)

	for i := m + n - 1; m > 0 && n > 0; i-- {
		if arr1[m-1] <= arr2[n-1] {
			arr1[i] = arr2[n-1]
			n--
		} else {
			arr1[i] = arr1[m-1]
			m--
		}
	}

	for ; n > 0; n-- {
		arr1[n-1] = arr2[n-1]
	}

	return arr1
}

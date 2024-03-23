package slide

func noRepeat(s string) int {
	if len(s) == 0 {
		return 0
	}

	freq := make(map[byte]int)
	left, right, res := 0, -1, 0

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]] += 1
			right++
		} else {
			res = max(res, right-left+1)
			freq[s[left]] -= 1
			left++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package slide

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	freqs := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freqs[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if freqs[toLower(s[i])] > 0 && freqs[toUpper(s[i])] > 0 {
			continue
		}
		left := longestNiceSubstring(s[:i])
		right := longestNiceSubstring(s[i+1:])
		if len(left) > len(right) {
			return left
		}
		return right
	}

	return s
}

func toUpper(b byte) byte {
	if b >= 97 {
		return b - 32
	}
	return b
}

func toLower(b byte) byte {
	if b >= 97 {
		return b
	}
	return b + 32
}

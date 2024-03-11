package dp

import "fmt"

/*
dp[i][j] = s[i] == s[j] && (j - i < 3 || dp[i+1][j-1])
*/
func longestPalindrome2(str string) string {
	if len(str) < 2 {
		return str
	}
	var dps = make([][]bool, len(str))
	for i := 0; i < len(str); i++ {
		dps[i] = make([]bool, len(str))
	}

	res := str[0:1]
	for i := len(str) - 2; i >= 0; i-- {
		for j := i; j < len(str); j++ {
			//fmt.Printf("i:%d j:%d\n", i, j)
			dps[i][j] = str[i] == str[j] && (j-i < 3 || dps[i+1][j-1])
			fmt.Printf("dp[%d][%d] = %t\n", i, j, dps[i][j])
			if dps[i][j] && j-i+1 > len(res) {
				res = str[i : j+1]
			}
		}
	}
	fmt.Println(dps)
	return res
}

func longestPalindrome(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		for j := len(str) - 1; j >= i; j-- {
			if isPalindrome(str[i:j+1]) && j+1-i > len(res) {
				res = str[i : j+1]
			}
		}
	}
	return res
}

func isPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	}

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

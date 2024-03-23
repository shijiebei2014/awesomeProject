package dp

import "fmt"

/*
爬楼梯=>一个数只能由1和2组成,问有几种组合
*/
func climbStairs(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func targetSum(nums []int, target int) int {
	var dp = make([]int, target+1)
	dp[0] = 1

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
		fmt.Printf("%d %v\n", num, dp)
	}

	return dp[target]
}

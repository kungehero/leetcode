/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
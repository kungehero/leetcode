/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		a, b = b, min(b+cost[i-1], a+cost[i-2])
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end


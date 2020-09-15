/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start

var count int

func findTargetSumWays(nums []int, S int) int {

	backTracking(nums, 0, 0, S)
	return count

}
func backTracking(nums []int, i, sum, s int) {
	if i == len(nums) {
		if sum == s {
			count++
		}
	} else {
		backTracking(nums, i+1, sum+nums[i], s)
		backTracking(nums, i+1, sum-nums[i], s)
	}
}

// @lc code=end


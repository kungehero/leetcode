/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package leetcode

// @lc code=start
func rob(nums []int) int {
	return tryRob(nums, 0)
}

func tryRob(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}
	res := 0
	for i := index; i < len(nums); i++ {
		res = max(res, nums[i]+tryRob(nums, i))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// @lc code=end


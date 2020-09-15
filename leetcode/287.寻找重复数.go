/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	slow = nums[slow]
	fast = nums[nums[fast]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	pre := 0
	pre1 := slow
	for pre1 != pre {
		pre = nums[pre]
		pre1 = nums[pre1]
	}
	return pre
}

// @lc code=end


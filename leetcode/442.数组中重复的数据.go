/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	var res []int
	if len(nums) < 1 {
		return res
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			res = append(res, nums[i])
		}
	}
	return res
}

// @lc code=end


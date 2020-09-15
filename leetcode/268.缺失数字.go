/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] == 0 && len(nums) == 1 {
		return 1
	}

	if nums[0] > 0 && len(nums) >= 1 {
		return nums[0] - 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			return nums[i-1] + 1
		}
	}
	return nums[len(nums)-1] + 1
}

// @lc code=end


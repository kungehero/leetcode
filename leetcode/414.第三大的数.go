/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	sort.Ints(nums)
	count := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			count++
		}
		if count == 3 {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// @lc code=end


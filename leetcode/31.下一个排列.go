/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	for i+1 < k {
		nums[i+1], nums[k] = nums[k], nums[i+1]
		i++
		k--
	}
}

// @lc code=end


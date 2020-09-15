/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}

	for i := 0; i < length; i++ {
		num := abs(nums[i])
		if num <= length {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[j]-nums[i]) <= t && abs(j-i) <= k {
				return true
			}
		}
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {

	maxNum := max(height)

	l, r := 0, 0
	for i := 1; i <= maxNum; i++ {
		for j := 0; j < len(height); j++ {
			if height[j] == i {

			}
		}
	}

}

// @lc code=end


/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	res := 0
	if len(height) < 1 {
		return res
	}
	l, r := 0, len(height)-1
	for l < r {
		h := min(height[l], height[r])
		res = max(res, (r-l)*h)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end


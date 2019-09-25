/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
func maxArea(height []int) int {
	res, front, rear := 0, 0, len(height)-1
	for front != rear {
		max := min(height[rear], height[front]) * (rear - front)
		if max > res {
			res = max
		}
		if height[rear] > height[front] {
			front++
		} else {
			rear--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

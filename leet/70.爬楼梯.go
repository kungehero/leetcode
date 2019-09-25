/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 next = previous + ‘previous of previous’ Exempli Gratia: input 3 = input(2) + input(1) = 2 + 1 = 3 input 4 = input(3) + input(2) = 3 + 2 = 5 input 5 = input(4) + input(3) = 5 + 3 = 8
*/
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	front, rear := 2, 3
	for i := 5; i <= n; i++ {
		cur := front + rear
		front = rear
		rear = cur
	}
	return front + rear
}

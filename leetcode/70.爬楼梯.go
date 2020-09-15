/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package leetcode

// @lc code=start
func climbStairs(n int) int {
	mem := make([]int, 100)
	mem[0] = 1
	mem[1] = 1
	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}

// @lc code=end

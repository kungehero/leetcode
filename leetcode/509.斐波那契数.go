/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	if N < 1 {
		return 0
	}
	mem := make([]int, 100)
	return helper(mem, N)
}

func helper(mem []int, N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	if mem[N] != 0 {
		return mem[N]
	}
	mem[N] = helper(mem, N-1) + helper(mem, N-2)
	return mem[N]
}

// @lc code=end

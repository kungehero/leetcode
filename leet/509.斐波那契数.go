/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

func fib(n int) int {
	fibs := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}
	return fibs[n]
}

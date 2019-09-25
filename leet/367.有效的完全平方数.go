/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */
func isPerfectSquare(n int) bool {
	var a int = 1
	for a*a < n {
		a++
	}
	return a*a == n
}

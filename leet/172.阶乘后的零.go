/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */
func trailingZeroes(n int) int {
	res := 0
	i := 5
	for n/i != 0 {
		res += n / i
		i *= 5
	}
	return res
}

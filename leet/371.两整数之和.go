/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */
func getSum(a int, b int) int {
	sum, co := a^b, a&b
	for co != 0 {
		co = co << 1
		sum, co = sum^co, sum&co
	}
	return sum
}

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	if r := num % 9; r == 0 {
		return 9
	} else {
		return r
	}
	return num
}

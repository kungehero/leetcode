/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */
func findComplement(num int) int {
	res := 0
	for res < num {
		res <<= 1
		res += 1
	}
	return num ^ res
}

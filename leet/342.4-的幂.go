/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */
func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	for num%4 == 0 {
		num /= 4
	}
	return num == 1
}

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n/2 != 0 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

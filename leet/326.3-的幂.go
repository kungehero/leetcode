/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */
func isPowerOfThree(n int) bool {
	if n == 2 || n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

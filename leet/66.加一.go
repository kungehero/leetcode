/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	a := make([]int, len(digits)+1)
	a[0] = 1
	return a
}

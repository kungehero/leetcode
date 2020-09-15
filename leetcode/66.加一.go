/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	r := len(digits) - 1
	for r >= 0 {
		if digits[r] == 9 {
			digits[r] = 0
		} else {
			if digits[r] < 9 {
				digits[r]++
				break
			}
		}
		r--
	}
	res := []int{}
	if digits[0] == 0 {
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}
	return digits
}

// @lc code=end


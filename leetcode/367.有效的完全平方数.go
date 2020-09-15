/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}
	for i := 1; i <= num/2; i++ {
		if num == i*i {
			return true
		}

		if num < i*i {
			return false
		}
	}
	return false
}

// @lc code=end


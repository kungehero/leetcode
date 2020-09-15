/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	for i := 1; n >= 0; i++ {
		if n >= i {
			n -= i
		} else {
			return i - 1
		}
	}
	return 0

}

// @lc code=end


/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {

	var res int
	count := 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			count = 0
		}
		res = max(res, count)

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


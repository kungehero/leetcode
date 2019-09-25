/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */
func findMaxConsecutiveOnes(nums []int) int {
	i, j := 0, 0
	for _, v := range nums {
		if v == 1 {
			i++
		} else {
			i = 0
		}
		j = max(j, i)
	}

	return j
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

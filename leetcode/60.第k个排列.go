/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	var used = make([]bool, n)
	str := ""
	backTrack(n, k, str)
	return str
}

func backTrack(n, k int, str string) {
	start := 0
	if len(str) == n {
		start++
		if str == k {
			return
		}

	}

	for i := 1; i <= n; i++ {
		str += string(i)

	}
}

// @lc code=end


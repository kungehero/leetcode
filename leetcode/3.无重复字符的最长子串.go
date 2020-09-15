/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	fre := [256]int{}
	l, r := 0, -1
	res := 0
	for l < len(s) {
		if r < len(s)-1 && fre[s[r+1]] == 0 {
			r++
			fre[s[r]]++
		} else {
			fre[s[l]]--
			l++
		}
		res = Max(res, r-l+1)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

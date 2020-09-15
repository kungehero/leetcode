/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	res := 0
	l, r := 0, 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if Palindrome(string(s[i:j])) {
				if res <= len(s[i:j]) {
					l, r = i, j
					res = len(s[i:j])
				}
			}
		}
	}
	return s[l:r]
}
func Palindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// @lc code=end


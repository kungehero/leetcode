/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		for !isNum(s[l]) && l < r {
			l++
		}
		for !isNum(s[r]) && l < r {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isNum(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z' || s >= '0' && s <= '9')
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	str := []byte(s)
	for l < r {
		for !isVowels(str[l]) && l < r {
			l++
		}
		for !isVowels(str[r]) && l < r {
			r--
		}
		if isVowels(s[l]) && isVowels(s[r]) {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}
	}
	return string(str)
}

func isVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}

// @lc code=end


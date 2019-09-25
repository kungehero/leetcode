/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	i, j := 0, 0
	var res, arr string

	for i < len(s)-1 {
		j++
		res += string(s[j-1])

		if Palindrome(res) {
			if len(arr) <= len(s[i:j]) {
				arr = s[i:j]
			}
		}
		if j == len(s) && i < len(s)-1 {
			i++
			j = i
			res = ""
		}
	}
	return arr
}

func Palindrome(s string) bool {
	for i, j := 0, len(s)-1; j > i; i, j = i+1, j-1 {
		if string(s[i]) != string(s[j]) {
			return false
		}
	}
	return true
}

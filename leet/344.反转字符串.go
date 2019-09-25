/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; j > i; i, j = i+1, j-1 {
		if s[i] != s[j] {
			s[i], s[j] = s[j], s[i]
		}
	}
}

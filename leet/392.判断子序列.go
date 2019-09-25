/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */
func isSubsequence(s string, t string) bool {
	a := 0
	l := len(s)
	if l == 0 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if s[a] == t[i] {
			a++
		}
		if a == l {
			return true
		}
	}
	return false
}

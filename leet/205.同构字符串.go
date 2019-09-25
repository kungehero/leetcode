/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := make(map[rune]int)
	b := make(map[rune]int)
	for k, v := range s {
		if _, ok := a[v]; ok {
			continue
		}
		a[v] = k
	}
	for k, v := range t {
		if _, ok := b[v]; ok {
			continue
		}
		b[v] = k
	}
	flag := true
	for i := 0; i < len(s); i++ {
		if a[rune(s[i])] != b[rune(t[i])] {
			flag = false
			break
		}
	}
	return flag
}

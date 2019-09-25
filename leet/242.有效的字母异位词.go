/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v] = 1
	}
	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v] -= 1
			continue
		}
		m[v] -= 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

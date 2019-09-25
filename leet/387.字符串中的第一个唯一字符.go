/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v] = 0
	}

	for k, v := range s {
		if m[v] == 0 {
			return k
		}
	}
	return -1
}

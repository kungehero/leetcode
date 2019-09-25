/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) == 0 {
		return res
	}
	left, right := 0, 0
	matchCount := 0
	m := make(map[rune]int)
	win := make(map[rune]int)
	for _, v := range p {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v]++
	}

	for right < len(s) {
		a := rune(s[right])
		if _, ok := m[a]; ok {
			win[a]++
			if win[a] == m[a] {
				matchCount++
			}
		}
		right++
		for matchCount == len(m) {
			if right-left == len(p) {
				res = append(res, left)
			}
			b := rune(s[left])
			if _, ok := m[b]; ok {
				win[b]--
				if win[b] < m[b] {
					matchCount--
				}
			}
			left++
		}
	}
	return res
}

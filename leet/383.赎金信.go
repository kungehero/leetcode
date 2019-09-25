/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 && len(magazine) == 0 {
		return true
	}
	if len(ransomNote) == 0 && len(magazine) != 0 {
		return true
	}
	if len(ransomNote) != 0 && len(magazine) == 0 {
		return false
	}
	m := make(map[rune]int)
	for _, v := range ransomNote {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v] = 1
	}

	for _, v := range magazine {
		if _, ok := m[v]; ok && m[v] != 0 {
			m[v]--
		}
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum == 0
}

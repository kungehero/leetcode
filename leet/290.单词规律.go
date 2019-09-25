import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */
func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")
	if len(pattern) != len(s) {
		return false
	}
	m := make(map[string]string)
	for i := 0; i < len(pattern); i++ {
		if _, ok := m[string(pattern[i])]; !ok {
			if isMapValue(m, s[i]) {
				return false
			}
			m[string(pattern[i])] = s[i]
		} else {
			if m[string(pattern[i])] != s[i] {
				return false
			}
		}
	}
	return true
}

func isMapValue(m map[string]string, value string) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

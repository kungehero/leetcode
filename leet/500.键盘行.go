import "strings"

/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */
func findWords(words []string) []string {
	m1 := map[string]bool{
		"q": true,
		"w": true,
		"e": true,
		"r": true,
		"t": true,
		"y": true,
		"u": true,
		"i": true,
		"o": true,
		"p": true,
	}
	m2 := map[string]bool{
		"a": true,
		"s": true,
		"d": true,
		"f": true,
		"g": true,
		"h": true,
		"j": true,
		"k": true,
		"l": true,
	}
	m3 := map[string]bool{
		"z": true,
		"x": true,
		"c": true,
		"v": true,
		"b": true,
		"n": true,
		"m": true,
	}
	var res []string
	for _, v := range words {
		f1, f2, f3 := true, true, true
		for i := 0; i < len(v); i++ {
			if _, ok := m1[strings.ToLower(string(v[i]))]; !ok {
				f1 = false
				break
			}
		}
		for i := 0; i < len(v); i++ {
			if _, ok := m2[strings.ToLower(string(v[i]))]; !ok {
				f2 = false
				break
			}
		}
		for i := 0; i < len(v); i++ {
			if _, ok := m3[strings.ToLower(string(v[i]))]; !ok {
				f3 = false
				break
			}
		}
		if f1 == true {
			res = append(res, v)
		}
		if f2 == true {
			res = append(res, v)
		}
		if f3 == true {
			res = append(res, v)
		}
	}
	return res
}

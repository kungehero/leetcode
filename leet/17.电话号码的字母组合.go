/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
func letterCombinations(digits string) []string {
	m := map[byte][]string{
		'0': []string{"0"},
		'1': []string{"1"},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	var res []string
	for i := 0; i < len(digits); i++ {
		v := digits[i]
		reslen := len(res)
		if reslen == 0 {
			res = m[v]
			continue
		}
		for j := 0; j < reslen; j++ {
			for _, value := range m[v] {
				res = append(res, res[j]+value)
			}
		}
		res = res[reslen:]
	}
	return res
}

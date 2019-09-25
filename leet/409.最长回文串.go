import "fmt"

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */
func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}
	res := 0
	m := make(map[rune]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v] = 1
	}
	fmt.Println(m)
	for _, v := range m {
		if v > 1 && v%2 == 0 {
			res += v
		} else {
			res += v - 1
		}
	}
	if len(s)-res == 0 {
		return res
	}
	return res + 1
}

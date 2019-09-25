/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)+1-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

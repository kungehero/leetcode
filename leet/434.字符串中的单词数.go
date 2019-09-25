/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */
func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	s += " "
	segmentCount := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] != ' ' && s[i] == ' ' {
			segmentCount++
		}
	}
	return segmentCount
}

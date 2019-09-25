/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
func lengthOfLastWord(s string) int {
	black := true
	i, j := 0, 0
	for k, l := 0, len(s); k < l; k++ {
		if string(s[k]) != " " {
			if black {
				i = k
				black = false
			}
			j = k + 1
		} else {
			black = true
		}
	}
	return j - i
}

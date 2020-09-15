/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	l, r := 0, len(s)-1
	flag := true
	for r >= 0 {
		if s[r] != ' ' {
			l++
			flag = false
		} else {
			if !flag {
				break
			}
		}
		r--
	}
	return l
}

// @lc code=end


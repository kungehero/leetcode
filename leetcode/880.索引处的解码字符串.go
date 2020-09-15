/*
 * @lc app=leetcode.cn id=880 lang=golang
 *
 * [880] 索引处的解码字符串
 */

// @lc code=start
func decodeAtIndex(S string, K int) string {
	if K == 0 {
		return ""
	}
	size := 0
	for _, c := range S {
		s := string(c)
		if unicode.IsDigit(c) {
			num, _ := strconv.Atoi(s)
			size *= num
		} else {
			size++
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		c := S[i]
		crune := rune(c)
		s := string(c)
		K %= size
		if unicode.IsLetter(crune) {
			if K == 0 {
				return s
			}
			size--
		}
		if unicode.IsDigit(crune) {
			num, _ := strconv.Atoi(s)
			size /= num
		}
	}
	return ""
}

// @lc code=end


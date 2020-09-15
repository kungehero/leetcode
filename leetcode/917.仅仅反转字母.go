/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	res := []byte(S)

	l, r := 0, len(S)-1
	for l < r {
		if (!isCheck(res[l]) && isCheck(res[r])) && l < r {
			l++
		} else if (isCheck(res[l]) && !isCheck(res[r])) && l < r {
			r--
		} else if isCheck(res[l]) && isCheck(res[r]) {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		} else {
			l++
			r--
		}
	}
	return string(res)
}

func isCheck(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}

// @lc code=end


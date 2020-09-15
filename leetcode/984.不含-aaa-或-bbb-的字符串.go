/*
 * @lc app=leetcode.cn id=984 lang=golang
 *
 * [984] 不含 AAA 或 BBB 的字符串
 */

// @lc code=start
func strWithout3a3b(A int, B int) string {
	str := ""
	i := 1
	if A > B {
		for A > 0 {
			if i <= 2 {
				if B >= 0 {
					str += "a"
					i++
				}
				A--
			} else {
				str += "b"
				i = 1
				B--
			}

		}
	} else {
		for B >= 0 {
			if i <= 2 {
				if A > 0 {
					str += "b"
					i++
				}
				B--
			} else {
				str += "a"
				i = 1
				A--
			}
		}
	}
	return str
}

// @lc code=end


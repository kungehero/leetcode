/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	res := []byte(s)
	var (
		stack []int
		str   string
	)

	for i := 0; i < len(s); i++ {
		if res[i] == '(' {
			stack = append(stack, i)
		} else if res[i] == ')' {
			l := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			reverse(res, l, i)
		}
	}

	for _, v := range res {
		if v != '(' && v != ')' {
			str += string(v)
		}

	}
	return str
}

func reverse(s []byte, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// @lc code=end


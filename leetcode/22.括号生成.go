/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := new([]string)
	str := []byte{}
	backtrack(res, str, 0, 0, n)
	return *res
}

func backtrack(res *[]string, str []byte, open, close, n int) {

	if n*2 == len(str) {
		/* 	tem := make([]byte, len(str))
		copy(str, tem) */
		*res = append(*res, string(str))
		return
	}
	if open < n {
		str = append(str, '(')
		backtrack(res, str, open+1, close, n)
		str = str[:len(str)-1]
	}
	if close < open {
		str = append(str, ')')
		backtrack(res, str, open, close+1, n)
		str = str[:len(str)-1]
	}
}

// @lc code=end


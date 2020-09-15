/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	res := [][]string{}
	path := []byte{}
	if len(s) < 1 {
		return res
	}
	for i := 0; i <= len(s); i++ {
		backTracking(s, res, path, 0, i)
	}

	return res
}

func backTracking(s string, res *[][]string, path []byte, k, start int) {
	if len(path) == k {
		tem := make([]byte, len(path))
		copy(tem, path)
		if valid(string(tem)) {
			*res = append(*res, string(tem))
		}
	}
	for i := start; i < len(s); i++ {
		path = append(path, s[i])
		backTracking(s, res, path, k, start)
		path = path[:len(path)-1]
	}

}
func valid(s []string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end


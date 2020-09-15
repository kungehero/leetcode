/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, d []string) string {
	mps := make(map[byte]int)
	var res string
	for _, v := range []byte(s) {
		mps[v]++
	}

	for _, s := range d {
		mpdinx := make(map[byte]int)
		for _, v := range []byte(s) {
			mpdinx[v]++
		}
		if isCheck(mps, mpdinx) {
			res = max(res, s)
		}
	}
	return res
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}

	if len(a) == len(b) {
		i := 0
		for i < len(a) {
			if a[i] > b[i] {
				return b
			}
			i++
		}
		return a
	}
	return b
}

func isCheck(mps, mpdinx map[byte]int) bool {
	for k, v := range mpdinx {
		if mps[k] < v {
			return false
		}
	}
	return true
}

// @lc code=end


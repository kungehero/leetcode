/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	mps := make(map[byte]int)
	mpdinx := make(map[byte]int)

	for _, v := range []byte(t) {
		mps[v]++
	}

	for _, v := range []byte(s) {
		mpdinx[v]++
	}

	return isCheck(mps, mpdinx)
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


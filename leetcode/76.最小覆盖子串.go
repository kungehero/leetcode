/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	mt := make(map[byte]int)
	ms := make(map[byte]int)
	length := math.MaxInt32

	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}

	resl, resr, l, r := 0, 0, 0, 0

	for ; r < len(s); r++ {
		if r < len(s) && mt[s[r]] > 0 {
			ms[s[r]]++
		}
		for check(mt, ms) && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				resl, resr = l, l+length
			}

			if _, ok := mt[s[l]]; ok {
				ms[s[l]] -= 1
			}
			l++
		}
	}
	return s[resl:resr]
}

func check(mt map[byte]int, ms map[byte]int) bool {
	for k, v := range mt {
		if ms[k] < v {
			return false
		}
	}
	return true
}

// @lc code=end


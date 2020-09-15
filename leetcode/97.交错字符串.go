/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n][m]
}

/* mp3 := make(map[string]int)
	mp12 := make(map[string]int)

	for _, v := range []byte(s3) {
		mp3[string(v)]++
	}

	for _, v := range []byte(s1 + s2) {
		mp12[string(v)]++
	}

	for k, v := range mp3 {
		if mp12[k] < v || mp12[k] > v {
			return false
		}
	}
	return true
} */

// @lc code=end


/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	if s == "" || len(s) == 0 {
		return 0
	}

	hash := make([]int, 26)
	l, maxCount, result := 0, 0, 0
	for r := 0; r < len(s); r++ {
		hash[s[r]-'A']++

		if maxCount < hash[s[r]-'A'] {
			maxCount = hash[s[r]-'A']
		}

		for r-l+1-maxCount > k {
			hash[s[l]-'A']--
			l++
		}

		if result < r-l+1 {
			result = r - l + 1
		}
	}
	return result
}

// @lc code=end


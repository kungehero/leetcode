/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	C := x ^ y
	count := 0

	s := fmt.Sprintf("%b", C)

	for i := range s {
		if s[i] == 49 {
			count++
		}
	}

	return count
}

// @lc code=end


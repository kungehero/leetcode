/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var res [][]int
	if rowIndex == 0 {
		return []int{1}
	}

	res = append(res, []int{1})
	count := 0

	for i := 1; i <= 33; i++ {
		m := []int{0}
		m = append(m, res[i-1]...)
		for j := 0; j < len(m)-1; j++ {
			m[j] = m[j] + m[j+1]
		}
		count++

		if count == rowIndex {
			return m
		}
		res = append(res, m)
	}
	return []int{}
}

// @lc code=end


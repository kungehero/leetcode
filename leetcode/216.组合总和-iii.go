/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var (
		res  [][]int
		path []int
	)
	backTracking(k, n, 1, &res, path)
	return res
}

func backTracking(k, n, start int, res *[][]int, path []int) {
	if k == 0 && n == 0 {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
		return
	}
	if k == 0 || n < 0 {
		return
	}
	for i := start; i <= 9; i++ {
		path = append(path, i)
		backTracking(k-1, n-i, i+1, res, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end


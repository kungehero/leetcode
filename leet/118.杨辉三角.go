/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */
func generate(numRows int) [][]int {
	var res [][]int

	if numRows == 0 {
		return res
	}
	for i := 1; i <= numRows; i++ {
		var path []int
		if i == 1 {
			path = append(path, 1)
		} else if i == 2 {
			path = append(path, 1, 1)
		} else {
			path = append(path, 1)
			for j := 1; j < i-1; j++ {
				path = append(path, res[i-2][j-1]+res[i-2][j])
			}
			path = append(path, 1)
		}
		res = append(res, path)
	}
	return res
}

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */
// @lc code=start
func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }
    if len(triangle) == 1 {
        return triangle[0][0]
    }
    for i := len(triangle)-2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            min := triangle[i+1][j]
            if triangle[i+1][j+1] < min {
                min = triangle[i+1][j+1]
            }
            triangle[i][j] = triangle[i][j] + min
        }
    }
    return triangle[0][0]
}
// @lc code=end


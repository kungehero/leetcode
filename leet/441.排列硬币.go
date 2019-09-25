import "math"

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */
func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(1+8*n)) - float64(1)) / 2)
}

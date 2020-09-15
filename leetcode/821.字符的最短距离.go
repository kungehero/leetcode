/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(S string, C byte) []int {

	temp := []int{}
	res := []int{}
	for i, v := range S {
		if byte(v) == C {
			temp = append(temp, i)
		}
	}

	for i := 0; i < len(S); i++ {
		minV := math.MaxInt32
		for j := 0; j < len(temp); j++ {
			minV = min(minV, abs(temp[j]-i))

		}
		res = append(res, minV)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	var res int
	m := make(map[int]int)
	for _, c := range C {
		for _, d := range D {
			m[c+d]++
		}
	}

	for _, a := range A {
		for _, b := range B {
			if val, ok := m[0-a-b]; ok {
				res += val
			}
		}
	}
	return res
}

// @lc code=end


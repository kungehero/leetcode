/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */
func missingNumber(nums []int) int {
	m := make(map[int]int)
	res := 0

	for _, v := range nums {
		m[v] = 0
	}
	for i := 0; i < len(nums)+1; i++ {
		if _, ok := m[i]; !ok {
			res = i
			break
		}
	}
	return res
}

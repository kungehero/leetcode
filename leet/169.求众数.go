/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 */
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] >= len(nums)/2 {
			return v
		}
		if _, ok := m[v]; ok {
			m[v] += 1
			continue
		}
		m[v] = 1
	}
	return 0
}

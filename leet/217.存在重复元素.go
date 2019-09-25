/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = 0
	}
	return false
}

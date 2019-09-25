/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, k := range nums {
		if j, ok := m[target-k]; ok {
			return []int{j, i}
		}
		m[k] = i
	}
	return nil
}

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	var res []int
	for i := 0; i < len(numbers); i++ {
		if v, ok := m[target-numbers[i]]; ok {
			if i >= v {
				res = append(res, v+1, i+1)
				return res
			}
		}
		m[numbers[i]] = i
	}
	return res
}

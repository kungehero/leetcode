/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := []int{}
	for i, v := range nums {
		if value, ok := m[target-v]; ok {
			res = append(res, value, i)
		}
		m[v] = i
	}
	return res
}

// @lc code=end


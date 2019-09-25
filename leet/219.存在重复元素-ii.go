/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for j, v := range nums {
		if i, ok := m[v]; ok && j-i <= k && nums[i] == nums[j] {
			return true
		}
		m[v] = j
	}
	return false
}

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for key, v := range nums {
		if _, ok := mp[v]; ok {
			if key-mp[v] <= k {
				return true
			}
		}
		mp[v] = key
	}
	return false
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			return true
		}
		mp[v]++
	}
	return false
}

// @lc code=end


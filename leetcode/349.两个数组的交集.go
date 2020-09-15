/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	mp := make(map[int]int)

	for _, v := range nums1 {
		mp[v]++
	}

	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			res = append(res, v)
			delete(mp, v)
		}
	}
	return res
}

// @lc code=end


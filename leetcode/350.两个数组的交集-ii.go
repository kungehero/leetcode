/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {

	var res []int
	mp := make(map[int]int)

	for _, v := range nums1 {
		mp[v]++
	}

	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			if mp[v] >= 1 {
				mp[v]--
				res = append(res, v)
			} else {
				return res
				//delete(mp, v)
			}

		}
	}
	return res
}

// @lc code=end


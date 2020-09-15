/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := (m + n) - 1
	for k >= 0 {
		if n == 0 {
			break
		}
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[k] = nums1[m-1]
			m--
		} else {
			nums1[k] = nums2[n-1]
			n--
		}
		k--
	}
}

// @lc code=end


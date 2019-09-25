/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	frontlen := m - 1
	rearlen := n - 1
	finallen := m + n - 1

	for rearlen >= 0 {
		if frontlen < 0 {
			nums1[finallen] = nums2[rearlen]
			rearlen--
		} else if nums2[rearlen] >= nums1[frontlen] {
			nums1[finallen] = nums2[rearlen]
			rearlen--
		} else {
			nums1[finallen] = nums1[frontlen]
			frontlen--
		}
		finallen--
	}
}

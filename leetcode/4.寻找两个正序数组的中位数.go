/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2

	} else {
		return float64(nums1[len(nums1)/2])
	}
	return 0
}

// @lc code=end


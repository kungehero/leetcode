/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}
	var cur []int
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			cur = append(cur, nums1[i])
			i++
		}
		for i < len(nums1) {
			if nums1[i] > nums2[j] {
				cur = append(cur, nums2[j])
				j++
			}
			break
		}
	}

	for ; len(nums1)-i > 0; i++ {
		cur = append(cur, nums1[i])
	}
	for ; len(nums2)-j > 0; j++ {
		cur = append(cur, nums2[j])
	}
	if len(cur)%2 == 0 {
		m := len(cur) / 2
		return float64(cur[m-1]+cur[m]) / 2
	} else {
		m := int(float64(len(cur) / 2))
		return float64(cur[m])
	}
}

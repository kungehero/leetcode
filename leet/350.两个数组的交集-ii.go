/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */
func intersect(nums1 []int, nums2 []int) []int {
	var arr []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return arr
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v] = 1
	}

	for _, v := range nums2 {
		if _, ok := m[v]; ok && m[v] != 0 {
			arr = append(arr, v)
			m[v] -= 1
		}
	}
	return arr
}

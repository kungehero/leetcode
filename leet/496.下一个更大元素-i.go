/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var res []int
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i
	}
	l2 := len(nums2) - 1
	for _, v := range nums1 {
		a := m[v]
		if a < l2 {
			for a < l2 {
				a++
				if v < nums2[a] {
					res = append(res, nums2[a])
					break
				} else if a == l2 && v > nums2[a] {
					res = append(res, -1)
				}
			}
		} else {
			res = append(res, -1)
		}
	}
	return res
}

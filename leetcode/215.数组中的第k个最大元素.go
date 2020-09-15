/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	nums = mergeSort(nums)

	return nums[len(nums)-k]
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	a := mergeSort(nums[:mid])
	b := mergeSort(nums[mid:])
	return merge(a, b)
}

func merge(a []int, b []int) (res []int) {
	l, r := 0, 0
	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			res = append(res, a[l])
			l++
		} else {
			res = append(res, b[r])
			r++
		}
	}
	res = append(res, a[l:]...)
	res = append(res, b[r:]...)
	return res
}

// @lc code=end


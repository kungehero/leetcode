/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	var sum int
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}
func quickSort(nums []int, left, right int) {

	if left >= right {
		return
	}
	explodindex := left

	for i := left + 1; i <= right; i++ {
		if nums[i] < nums[left] {
			explodindex++
			nums[i], nums[explodindex] = nums[explodindex], nums[i]
		}
	}
	nums[left], nums[explodindex] = nums[explodindex], nums[left]
	quickSort(nums, left, explodindex-1)
	quickSort(nums, explodindex+1, right)
}

// @lc code=end


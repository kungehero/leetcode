/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	outLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[outLen] = nums[i]
			outLen++
		}
	}
	return outLen
}

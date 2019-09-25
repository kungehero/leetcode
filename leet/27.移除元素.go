/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	outLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val && i < len(nums) {
			nums[outLen] = nums[i]
			outLen++
		}
	}
	return outLen
}

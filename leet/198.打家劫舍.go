/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := make([]int, len(nums)+1)
	res[0] = 0
	res[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		res[i] = Max(res[i-1], res[i-2]+nums[i-1])
	}
	return res[len(nums)]
}
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

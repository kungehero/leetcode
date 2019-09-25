/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
		index++
	}
	return index
}

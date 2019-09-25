/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
func moveZeroes(nums []int) {
	zerocount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zerocount++
		} else {
			nums[i-zerocount] = nums[i]
		}
	}
	for i := len(nums) - zerocount; i < len(nums); i++ {
		nums[i] = 0
	}
}

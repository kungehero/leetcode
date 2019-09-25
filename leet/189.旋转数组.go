/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */
func rotate(nums []int, k int) {
	if len(nums) != 0 {
		copy(nums, append(nums[len(nums)-k%len(nums):], nums[0:len(nums)-k%len(nums)]...))
	}
}

//三次旋转
//使用栈
//3个for

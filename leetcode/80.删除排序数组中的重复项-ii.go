/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */
// @lc code=start
func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	k, count := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// @lc code=end

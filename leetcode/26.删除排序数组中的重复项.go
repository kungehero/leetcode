/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */
package leetcode

// @lc code=start
func removeDuplicates(nums []int) int {

	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

// @lc code=end

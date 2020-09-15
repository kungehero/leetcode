/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// @lc code=end


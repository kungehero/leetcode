/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	for i := 0; (i + k) <= len(nums); i++ {
		num := getMin(nums[i : i+k])
		res = append(res, num)
	}
	return res
}

func getMin(nums []int) int {
	i := 1
	max := nums[0]
	for i < len(nums) {
		if nums[i] > max {
			max = nums[i]
		}
		i++
	}
	return max
}

// @lc code=end


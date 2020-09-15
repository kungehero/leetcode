/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */
// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	l, r := 0, -1
	sum := 0
	res := len(nums) + 1

	for l < len(nums) {
		if r < len(nums)-1 && sum < s {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = Min(res, r-l+1)
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

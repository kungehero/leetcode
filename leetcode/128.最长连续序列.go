/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	res := 0
	l, r := 0, 1
	sort.Ints(nums)
	for r < len(nums) {
		if nums[r]-nums[r-1] == 1 {
			r++
		} else if nums[r] == nums[r-1] {
			l++
			r++
		} else {
			l = r
			r++
		}
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


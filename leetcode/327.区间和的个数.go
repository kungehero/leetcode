	/*
	 * @lc app=leetcode.cn id=327 lang=golang
	 *
	 * [327] 区间和的个数
	 */

	// @lc code=start
	func countRangeSum(nums []int, lower int, upper int) int {
		var count int
		if len(nums) == 0 {
			return 0
		}
		for i := 0; i < len(nums); i++ {
			for j := i; j < len(nums); j++ {
				sum := 0
				for k := i; k <= j; k++ {
					sum += nums[k]
				}
				if lower <= sum && sum <= upper {
					count++
				}
			}

		}
		return count
	}

	// @lc code=end


/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start
func countSmaller(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++

			}
		}
		res = append(res, count)
	}
	return res
}

// @lc code=end


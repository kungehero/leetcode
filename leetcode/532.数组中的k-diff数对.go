/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				res++
				break
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


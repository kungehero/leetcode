/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				numbers[0], numbers[1] = i+1, j+1
			}
		}
	}
	return numbers[:2]
}

// @lc code=end


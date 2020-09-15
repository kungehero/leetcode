/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	if nums[0] > 1 {
		for i := 1; i < nums[0]; i++ {
			res = append(res, i)
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			count := nums[i-1] + 1
			for count < nums[i] {
				res = append(res, count)
				count++
			}
			continue
		}
	}

	if nums[len(nums)-1] < len(nums) {
		num := nums[len(nums)-1] + 1
		for num <= len(nums) {
			res = append(res, num)
			num++
		}
	}

	return res
}

// @lc code=end


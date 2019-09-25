import "sort"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */
func thirdMax(nums []int) int {
	sort.Ints(nums)
	count := 1
	target := 0
	for i := len(nums) - 1; i > 0; i-- {
		if count != 3 {
			if nums[i] != nums[i-1] {
				count++
			}
		} else {
			target = i
			break
		}
	}
	if count < 3 {
		return nums[len(nums)-1]
	}
	return nums[target]
}

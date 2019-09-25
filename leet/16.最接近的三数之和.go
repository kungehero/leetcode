import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */
func threeSumClosest(nums []int, target int) int {
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		v := i + 1
		left, right, target := 0, len(nums)-2, -nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if target == sum+nums[i]+v || target == sum+nums[i]-v {
				count = nums[i] + nums[left] + nums[right]
				break
			} else if target > sum+nums[i] {
				left++
			} else if target < sum+nums[i] {
				right--
			}
		}
	}
	return count
}

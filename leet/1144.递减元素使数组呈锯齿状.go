/*
 * @lc app=leetcode.cn id=1144 lang=golang
 *
 * [1144] 递减元素使数组呈锯齿状
 */
//A modern and intuitive terminal-based text editor
func movesToMakeZigzag(nums []int) int {
	return min(move(0, nums), move(1, nums))
}

//start=0 偶数  start=1 奇数
func move(start int, nums []int) int {
	res := 0
	for i := start; i < len(nums); i += 2 {
		if i == 0 && nums[i] > nums[i+1] {
			res += nums[i] - nums[i+1] + 1
		} else if i > 0 && i < len(nums)-1 {
			m := min(nums[i-1], nums[i+1])
			if nums[i] >= m {
				res += nums[i] - m + 1
			}
		} else if i == len(nums)-1 && nums[i] > nums[i-1] {
			res += nums[i] - nums[i-1] + 1
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

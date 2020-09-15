/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var (
		res  [][]int
		path []int
	)
	backTracking(nums, 2, 4, &res, path)
	return res
}

func backTracking(nums []int, target, count int, res *[][]int, path []int) {
	if len(path) == 4 && target == 0 {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
		return
	}

	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		backTracking(nums, target-nums[i], 4, res, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end


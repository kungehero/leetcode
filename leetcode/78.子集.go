/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	for k := 0; k < len(nums)+1; k++ {
		backTrack(nums, 0, k, &res, path)
	}
	return res
}

func backTrack(nums []int, start, k int, res *[][]int, path []int) {
	if len(path) == k {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backTrack(nums, i+1, k, res, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end


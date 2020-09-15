/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	res := [][]int{}
	path := []int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)+1; i++ {
		backTraacking(nums, 0, i, &res, path)
	}
	return res
}

func backTraacking(nums []int, start, k int, res *[][]int, path []int) {
	if len(path) == k {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backTraacking(nums, i+1, k, res, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end


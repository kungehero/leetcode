/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	backtrack(nums, &res, []int{})
	return res
}

func backtrack(nums []int, res *[][]int, path []int) {

	if len(path) == len(nums) {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
		return
	}
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, res, path)
		path = path[:len(path)-1]
	}

}

// @lc code=end


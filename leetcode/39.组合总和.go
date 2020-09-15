/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	length := len(candidates)
	sort.Ints(candidates)

	dfs(candidates, length, target, 0, []int{}, &res)
	return res
}

func dfs(candidates []int, length int, target int, begin int, path []int, res *[][]int) {
	if target == 0 {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
		return
	}
	for i := begin; i <= length-1; i++ {
		if target-candidates[i] < 0 {
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, length, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}

// @lc code=end


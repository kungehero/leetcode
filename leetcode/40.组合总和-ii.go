/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	path := []int{}
	dfs(candidates, target, 0, path, &res)
	return res
}

func dfs(candidates []int, target int, begin int, path []int, res *[][]int) {
	if target == 0 {
		tem := make([]int, len(path))
		copy(tem, path)
		*res = append(*res, tem)
		return
	}

	// 小剪枝

	for i := begin; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		dfs(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]

	}
}

// @lc code=end


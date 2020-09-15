/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start

func combine(n int, k int) [][]int {
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	result := [][]int{}
	backtrack(1, []int{}, k, n, &result)
	return result
}

func backtrack(start int, path []int, k, n int, result *[][]int) {
	if len(path) == k {
		tem := make([]int, len(path))
		copy(tem, path)
		*result = append(*result, tem)
		return
	}
	for i := start; i < n+1; i++ {
		path = append(path, i)
		backtrack(i+1, path, k, n, result)
		path = path[:len(path)-1]
	}
}

// @lc code=end


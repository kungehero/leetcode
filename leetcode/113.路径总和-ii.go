/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	if root == nil {
		return [][]int{}
	}
	backTracking(root, target, &res, path)
}

func backTracking(root *TreeNode, target int, res *[][]int, path []int) {
	if root.Val == target {
		tem := make([]int, len(path))
		*res = append(*res, tem)
		return
	}
	if target-root.Val > 0 {
		path = append(path, root.Val)
		backTracking(root.Left, target-root.Val, res, path)
		backTracking(root.Right, target-root.Val, res, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
    func isSymmetric(root *TreeNode) bool {
		return dfs(root, root)
	}
	
	func dfs(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		return dfs(p.Left, q.Right) && p.Val == q.Val && dfs(p.Right, q.Left)
	}

// @lc code=end


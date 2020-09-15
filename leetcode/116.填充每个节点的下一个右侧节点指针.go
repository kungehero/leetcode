/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
		root.Right = nil
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

// @lc code=end


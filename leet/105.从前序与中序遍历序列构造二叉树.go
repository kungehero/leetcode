/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

func buildTreeHelper(preorder []int, p_start int, p_end int, inorder []int, i_start int, i_end int) *TreeNode {
	// preorder 为空，直接返回 null
	if p_start == p_end {
		return nil
	}
	//前序遍历的第一个值，根节点的值
	root_val := preorder[p_start]
	root := &TreeNode{Val: root_val}
	//在中序遍历中找到根节点的位置
	i_root_index := 0

	//取出根节点在inorder中的位置
	for i := i_start; i < i_end; i++ {
		if root_val == inorder[i] {
			i_root_index = i
			break
		}
	}

	leftNum := i_root_index - i_start
	//递归的构造左子树
	root.Left = buildTreeHelper(preorder, p_start+1, p_start+leftNum+1, inorder, i_start, i_root_index)
	//递归的构造右子树
	root.Right = buildTreeHelper(preorder, p_start+leftNum+1, p_end, inorder, i_root_index+1, i_end)
	return root
}

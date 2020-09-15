/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
 func postorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }

    postOrder(root,&res)
    return res
}

func postOrder(t *TreeNode,s *[]int) {
	
	if t.Left != nil {
		postOrder(t.Left,s)
	}

   
	if t.Right != nil {
		postOrder(t.Right,s)
	}

     *s = append(*s,t.Val)

}
// @lc code=end


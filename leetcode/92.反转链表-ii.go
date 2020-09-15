/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	res := &ListNode{Next: head}
	newhead := res
	for i := 1; i < m; i++ {
		newhead = newhead.Next
	}

	var pre *ListNode
	cur := newhead.Next

	for i := m; i <= n; i++ {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	newhead.Next.Next = cur
	newhead.Next = pre
	return res.Next
}

// @lc code=end


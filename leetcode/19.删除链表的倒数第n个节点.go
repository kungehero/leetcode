/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Next: head}
	slow, fast := res, res
	for n >= 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return res.Next
}

// @lc code=end


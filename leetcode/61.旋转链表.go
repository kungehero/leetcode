/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	res := &ListNode{Next: head}
	slow := res
	fast := res
	for head != nil {
		length++
		head = head.Next
	}
	m := 0
	if length == k {
		return head
	}
	if length-k > 0 {
		m = k
	} else {
		m = k % length
	}
	for m >= 0 {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newhead, tail := reverse(slow.Next)
	slow.Next = nil
	tail.Next = res.Next
	return newhead
}

func reverse(head *ListNode) (res *ListNode, tail *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre, head
}

// @lc code=end

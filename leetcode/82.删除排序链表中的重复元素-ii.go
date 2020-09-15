/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	a := dummyHead
	b := head
	for b != nil && b.Next != nil {
		if a.Next.Val == b.Next.Val {
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}
			a.Next = b.Next
			b = b.Next
		} else {
			a = a.Next
			b = b.Next
		}
	}
	return dummyHead.Next
}

// @lc code=end

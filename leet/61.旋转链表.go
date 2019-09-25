/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 To simplify, let’s assume k is less than a number of ListNode first.
Then we need to do is let a fast pointer move to the k th node first, then set a slow pointer to the head and move both slow and fast pointers until the fast pointer reach to the end.

fast is on the last node and slow is k th point before the fast one. So we can arrange the node like so that:

a new head is positioned at the next node of slow.
slow.Next should nil, means an end.
fast.Next should an original head.
So…remaining tricky part is when k is more than a length of the list. To handle this, I calculate the remaining rotation number and call another rotateRight.
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	newHead := slow.Next
	slow.Next, fast.Next = nil, head
	return newHead
}

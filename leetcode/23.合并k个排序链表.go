/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	var res *ListNode
	for _, v := range lists {
		if v != nil {
			res = mergeTwoLists(res, v)
		}
	}
	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return res.Next
}
// @lc code=end


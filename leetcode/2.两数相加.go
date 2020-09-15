/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	sum := 0
	cur := res
	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		fmt.Println(sum)
		res.Next = &ListNode{Val: sum % 10}
		res = res.Next
		sum = sum / 10
	}
	return cur.Next
}

// @lc code=end

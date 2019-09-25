/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成

 Consider a tree that each node has each "(" or ")" character. We manage how many more left parenthesis should appear (leftCount) and same value for right (rightCount). As long as leftCount > 0 , we should add "(" as a child node and whenever we add a left parenthesis, we also increment rightCount because the corresponding right one should appear later. Then as long as rightCount > 0, we can also add ")" as a child node.

I think the complexity is 2^n because a number of children is maximum 2 and the height of the tree equals to n.
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	var res []string
	backtrack(n, 0, []byte{}, &res)
	return res
}

func backtrack(leftCount int, rightCount int, curBytes []byte, res *[]string) {
	// base case
	if leftCount == 0 && rightCount == 0 {
		*res = append(*res, string(curBytes))
		return
	}

	curLen := len(curBytes)
	if leftCount > 0 {
		newBytes := make([]byte, curLen+1)
		copy(newBytes[:curLen], curBytes)
		newBytes[curLen] = '('

		backtrack(leftCount-1, rightCount+1, newBytes, res)
	}
	if rightCount > 0 {
		newBytes := make([]byte, curLen+1)
		copy(newBytes[:curLen], curBytes)
		newBytes[curLen] = ')'

		backtrack(leftCount, rightCount-1, newBytes, res)
	}
}

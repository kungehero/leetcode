/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	stack:=[]int{-1}
	res:=0
	for i := 0; i < len(s); i++ {
		if s[i]=='('{
			stack=append(stack,i)
		}else{
			stack=stack[:len(stack)-1]
			if len(stack)==0{
				stack=append(stack,i)
			}else{
				res=max(res,i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(a ,b int) int{
	if a>b{
		return a
	}
	return b
}
// @lc code=end


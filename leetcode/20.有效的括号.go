/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
m:=make(map[byte]byte)
m[')']='('
m[']']='['
m['}']='{'

stack:=[]byte{}

if len(s)<1{
	return true
}

for i:=0;i<len(s);i++{
	if s[i]=='('||s[i]=='['||s[i]=='{'{
		stack=append(stack,s[i])
	}else{
		if  len(stack)>0&& m[s[i]]==stack[len(stack)-1]{
			stack=stack[:len(stack)-1]
		}else{
			return false
		}
	}
}
return len(stack)==0
}
// @lc code=end

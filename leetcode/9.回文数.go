/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {

	if x<0 ||(x%10==0&&x>0){
		return false
	}

	reverseNumeber:=0

	for x>reverseNumeber{
		reverseNumeber=reverseNumeber*10+x%10
		x/=10
	}

	return x ==reverseNumeber ||x==reverseNumeber/10

}
// @lc code=end


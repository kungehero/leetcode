package leet

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 输入: ["flower","flow","flight"]
输出: "fl"
目标 过程分析 测试 结果 提交
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pv := len(strs[0])

	for {
		if pv == 0 {
			return ""
		}
		prefix := strs[0][:pv]
		flag := true
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) < pv || strs[i][:pv] != prefix {
				flag = false
				break
			}
		}
		if flag {
			return prefix
		} else {
			flag = false
			pv--
		}
	}
}

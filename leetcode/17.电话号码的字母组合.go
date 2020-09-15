/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return nil
	}
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	res := []string{""}
	for _, v := range digits {
		tem := []string{}
		for _, value := range m[v-'2'] {
			for _, v2 := range res {
				tem = append(tem, string(v2)+string(value))
			}
		}
		res = tem
	}
	return res
}

// @lc code=end


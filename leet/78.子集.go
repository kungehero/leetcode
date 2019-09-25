/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {
		var path [][]int

		for j := 0; j < len(res); j++ {
			var temp []int = res[j]
			temp = append(temp, nums[i])
			path = append(path, temp)
		}
		/* for _, v := range res {
			var temp []int = v
			temp = append(temp, nums[i])
			path = append(path, temp)
		} */
		res = append(res, path...)
	}
	return bn/
}

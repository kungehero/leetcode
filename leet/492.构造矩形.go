import "math"

/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */
func constructRectangle(area int) []int {
	var res []int
	a := int(math.Sqrt(float64(area)))
	for area%a != 0 {
		a--
	}
	res = append(res, area/a, a)
	return res
}

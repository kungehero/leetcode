/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
type NumArray struct {
	arr []int
	//length int
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums}
}

func (this *NumArray) Update(i int, val int) {
	this.arr[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	var res int
	for i <= j {
		res += this.arr[i]
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end


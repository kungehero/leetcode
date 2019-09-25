/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */
type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == j && j < len(this.arr)-1 {
		return this.arr[i]
	}
	sum := 0
	for ik := i; ik <= j; ik++ {
		sum += this.arr[ik]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

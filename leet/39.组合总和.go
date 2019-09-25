import "sort"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	return res
}


func findCombinationSum( residue int,  start int, path []int) {
	if (residue == 0) {
		res=append(res,path)
		return;
	}
	// 优化添加的代码2：在循环的时候做判断，尽量避免系统栈的深度
	// residue - candidates[i] 表示下一轮的剩余，如果下一轮的剩余都小于 0 ，就没有必要进行后面的循环了
	// 这一点基于原始数组是排序数组的前提，因为如果计算后面的剩余，只会越来越小
	for int i = start; i < len && residue - candidates[i] >= 0; i++ {
		path=append(path,candidates[i])
		// 【关键】因为元素可以重复使用，这里递归传递下去的是 i 而不是 i + 1
		findCombinationSum(residue - candidates[i], i, );
		 
	}
}


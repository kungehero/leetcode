/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(arr []int, l int, r int) int {
	if l >= r {
		return 0
	}
	count, mid := 0, (l+r)>>1
	count = mergeSort(arr, l, mid) + mergeSort(arr, mid+1, r) // 1
	temp, i, k := make([]int, r-l+1), l, 0
	for j, idx := mid+1, l; j <= r; j++ { // 2
		for ; i <= mid && arr[i] < arr[j]; i++ {
			temp[k], k = arr[i], k+1 // 2.1
		}
		temp[k], k = arr[j], k+1                      // 2.2
		for idx <= mid && (arr[idx]+1)>>1 <= arr[j] { // 3
			idx++ // 3.1
		}
		count += mid - idx + 1 // 3.2
	}
	copy(arr[l+k:], arr[i:mid+1]) // 2.3
	copy(arr[l:], temp[:k])       // 4
	return count
}

// @lc code=end


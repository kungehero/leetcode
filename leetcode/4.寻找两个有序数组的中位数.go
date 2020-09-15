/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    tem:=[]int{}
    tem=append(tem,nums1...)
    tem=append(tem,nums2...)
    var res float64
    flag:=0
    sort.Ints(tem)
    if len(tem)>=1{
       flag= len(tem)/2
    }

    if len(tem)%2==0{
        cur:=tem[flag-1]+tem[flag]
        res=float64(float64(cur)/2.0)
       
    }else{
        res=float64(tem[flag])
    }
    return res
    // write code here
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {

	if len(nums)==0{
		return []int{-1,-1}
	}
	if len(nums)==1{
		if nums[0]==target{
			return []int{0,0}
		}else{
		return []int{-1,-1}
		}
	}
	res:=[]int{}
	l,r:=0,len(nums)-1
	
	for l<=r{
		mid:=(l+r)/2
		if nums[mid]==target{
			j,k:=mid,mid
	for j-1>=0&&nums[j-1]==target{
		j--
	}
	for k+1<=len(nums)-1&&nums[k+1]==target{
		k++
	}
	res=append(res,j,k)
	  break
		}

		if nums[mid]<=target{
            l=mid+1
		}else{
            r=mid-1
		}
	}
	
if len(res)>0{
	return res
}
return []int{-1,-1}
}
// @lc code=end


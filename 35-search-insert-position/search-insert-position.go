func searchInsert(nums []int, target int) int {
      left,right:=0,len(nums)-1
      for left<=right{
        center:=left+(right-left)/2
        if nums[center]==target{
            return center
        }else if nums[center]<target{
            left=center+1
        }else{
            right=center-1
        }
      }
      return left
}
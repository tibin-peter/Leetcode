func maximumDifference(nums []int) int {
   maxDiff:=0
   flag:=false
   for i:=0;i<len(nums);i++{
    for j:=i+1;j<len(nums);j++{
        if nums[i]<nums[j]&&maxDiff<nums[j]-nums[i]{
            maxDiff=nums[j]-nums[i]
            flag=true
        }
    }
   } 
   if !flag{
    return -1
   }
   return maxDiff
}
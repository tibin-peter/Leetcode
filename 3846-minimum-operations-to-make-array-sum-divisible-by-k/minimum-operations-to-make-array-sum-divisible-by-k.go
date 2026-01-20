func minOperations(nums []int, k int) int {
   sum:=0
   for _,v:=range nums{
    sum+=v
   } 
   if sum%k==0{
    return 0
   }
   operations:=sum%k
   return operations
}
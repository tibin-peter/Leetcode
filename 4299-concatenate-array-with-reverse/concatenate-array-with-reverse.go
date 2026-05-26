func concatWithReverse(nums []int) []int {
    output:=make([]int,0,len(nums)*2)
    output=append(output,nums...)
    for i:=len(nums)-1;i>=0;i--{
        output=append(output,nums[i])
    }
  return output
}
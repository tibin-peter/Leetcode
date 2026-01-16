func findMaxAverage(nums []int, k int) float64 {
    sum:=0
    for i:=0;i<k;i++{
        sum+=nums[i]
    }
    maxsum:=sum
    for right:=k;right<len(nums);right++{
        sum+=nums[right]
        sum-=nums[right-k]
    if sum>maxsum{
        maxsum=sum
    }
    }
    return float64(maxsum)/float64(k)
}
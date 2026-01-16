func numOfSubarrays(arr []int, k int, threshold int) int {
    sum:=0
    count:=0
    for i:=0;i<k;i++{
        sum+=arr[i]
    }
    if sum/k>=threshold{
        count++
    }
    for right:=k;right<len(arr);right++{
        sum+=arr[right]
        sum-=arr[right-k]
        if sum/k>=threshold{
            count++
        }
    }
    return count
}
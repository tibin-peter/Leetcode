func getConcatenation(nums []int) []int {
    n:=len(nums)
    ans:=make([]int,2*n)
    ans=append(nums,nums...)
    return ans
}
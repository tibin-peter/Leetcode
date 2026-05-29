func minElement(nums []int) int {
    minElement:=getDigitSum(nums[0])
    for _,v:=range nums{
        sum:=getDigitSum(v)
        if sum<minElement{
            minElement=sum
        }
    }
    return minElement
}
func getDigitSum(val int)int{
    sum:=0
    for val>0{
        sum+=val%10
        val/=10
    }
    return sum
}
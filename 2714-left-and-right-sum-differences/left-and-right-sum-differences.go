func leftRightDifference(nums []int) []int {
    leftSum:=make([]int,len(nums))
    rightSum:=make([]int,len(nums))

    leftSumVal:=0
    rightSumVal:=0

    for i:=0;i<len(nums);i++{
        if i!=0{
            leftSumVal+=nums[i-1]
            leftSum[i]=leftSumVal
        }
    }
    for i:=len(nums)-1;i>=0;i--{
        if i<len(nums)-1{
            rightSumVal+=nums[i+1]
            rightSum[i]=rightSumVal
        }
    }

    for i:=0;i<len(nums);i++{
        dif:=leftSum[i]-rightSum[i]
        if dif<0{
            dif= -dif
        }
        nums[i]=dif
    }
    return nums
}
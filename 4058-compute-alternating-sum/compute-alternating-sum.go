func alternatingSum(nums []int) int {
    altSum:=0
    for i,v:=range nums{
        if i%2==0{
            altSum+=v
        }else{
            altSum-=v
        }
    }
    return altSum
}
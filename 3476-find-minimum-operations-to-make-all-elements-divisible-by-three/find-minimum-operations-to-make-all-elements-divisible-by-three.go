func minimumOperations(nums []int) int {
    operations:=0
    for _,v:=range nums{
        if v%3==1{
            operations++
        }
        if v%3==2{
            operations++
        }
    }
    return operations
}
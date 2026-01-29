func createTargetArray(nums []int, index []int) []int {
    target:=[]int{}
    for i:=0;i<len(nums);i++{
        pos:=index[i]
        target = append(target[:pos],append([]int{nums[i]}, target[pos:]...)...,)
    }
    return target
}
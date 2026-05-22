func containsNearbyDuplicate(nums []int, k int) bool {
    hash:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        if lastIndex,exist:=hash[nums[i]];exist{
            if i-lastIndex<=k{
                return true
            }
        }
        hash[nums[i]]=i
    }
    return false
}
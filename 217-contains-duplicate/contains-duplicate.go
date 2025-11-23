func containsDuplicate(nums []int) bool {
    mp:=make(map[int]int)
    for i,v:= range nums{
        if _,exists:=mp[v];exists{
            return true
        }
        mp[v]=i
    }
    return false
}
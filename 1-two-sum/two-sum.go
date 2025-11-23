func twoSum(nums []int, target int) []int {
    mp:=make(map[int]int)//storing value and index

    for i,v:=range nums{
        complement:=target-v

        if idx,exists:=mp[complement];exists{
            return []int{idx,i}
        }
        mp[v]=i
    }
    return []int{}//if no pair found
}
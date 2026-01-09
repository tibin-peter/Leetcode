func singleNumber(nums []int) int {
    hash:=make(map[int]int)
    var count int
    for _,v:= range nums{
        hash[v]++
    }
    for i,c:=range hash{
        if c==1{
            count=i
        }
    }
    return count
}
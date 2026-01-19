func smallerNumbersThanCurrent(nums []int) []int {
    sorted:=append([]int{},nums...)
    sort.Ints(sorted)
     mp:=map[int]int{}
     for i,v:=range sorted{
        if _,ok:=mp[v];!ok{
            mp[v]=i
        }
     }
     arr:=make([]int,len(nums))
     for i,v:=range nums{
        arr[i]=mp[v]
     }
     return arr
    }


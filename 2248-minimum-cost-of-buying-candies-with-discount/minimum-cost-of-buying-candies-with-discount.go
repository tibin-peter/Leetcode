func minimumCost(cost []int) int {
    sort.Slice(cost,func(a,b int)bool{
        return cost[a]>cost[b]
    })
    n:=len(cost)
    count:=0
    for i:=0;i<n;i+=3{// here i+3 work 
        count+=cost[i]// adding first one
        if i+1<n{
            count+=cost[i+1]// adding second one
        }
        //skip the third one
    }
    return count
}
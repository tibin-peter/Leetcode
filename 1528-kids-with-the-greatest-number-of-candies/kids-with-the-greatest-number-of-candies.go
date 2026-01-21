func kidsWithCandies(candies []int, extraCandies int) []bool {
    arr:=[]bool{}
    var greatest int
    for _,v:=range candies{
        if v>greatest{
            greatest=v
        }
    }
    for i:=0;i<len(candies);i++{
        total:=candies[i]+extraCandies
        if total>=greatest{
            arr=append(arr,true)
        }else{
            arr=append(arr,false)
        }
    }
    return arr
}
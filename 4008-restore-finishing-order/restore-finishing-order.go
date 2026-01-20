func recoverOrder(order []int, friends []int) []int {
    hash:=make(map[int]bool)
    finishingOrder:=[]int{}
    for _,v:=range friends{
        hash[v]=true
    }
    for _,v:=range order{
        if hash[v]{
            finishingOrder=append(finishingOrder,v)
        }
    }
    return finishingOrder
}
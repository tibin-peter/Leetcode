func minimumRecolors(blocks string, k int) int {
    count:=0
    for i:=0;i<k;i++{
        if blocks[i]=='W'{
            count++
        }
    }
    maxcount:=count
    for right:=k;right<len(blocks);right++{
        if blocks[right]=='W'{
            count++
        }
        if blocks[right-k]=='W'{
            count--
        }
        if maxcount>count{
            maxcount=count
        }
    }
    return maxcount
}
func isAnagram(s string, t string) bool {
    hash:=make(map[rune]int)

    if len(s)!=len(t){
        return false
    }

    for _,v:=range s{
        hash[v]++
    }
    for _,v:=range t{
        hash[v]--
        if hash[v]<0{
            return false
        }
    }
    return true
}
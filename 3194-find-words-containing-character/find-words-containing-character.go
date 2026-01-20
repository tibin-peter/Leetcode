func findWordsContaining(words []string, x byte) []int {
    var arr[] int
    for i:=0;i<len(words);i++{
        for _,v:=range words[i]{
            if byte(v)==x{
                arr=append(arr,i)
                break
            }
        }
    }
    return arr
}
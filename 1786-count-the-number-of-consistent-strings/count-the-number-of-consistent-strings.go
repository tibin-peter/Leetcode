func countConsistentStrings(allowed string, words []string) int {
    count:=0
    hashed:=make(map[rune]bool)
    for i:=0;i<len(allowed);i++{
        hashed[rune(allowed[i])]=true
    }
    for i:=0;i<len(words);i++{
        for _,v:=range words[i]{
            if !hashed[v]{
                count++
                break
            }
        }
    }
    return len(words)-count
}
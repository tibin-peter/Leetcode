func groupAnagrams(strs []string) [][]string {
    hashmap:=make(map[string][]string)//intialize the hashmap
    for _,s:=range strs{
        runes:=[]rune(s)//convert into rune
        sort.Slice(runes, func(i, j int) bool {//sort the rune for key
          return runes[i] < runes[j]
        })
        key:=string(runes)//convert backto string
        hashmap[key]=append(hashmap[key],s)//storing in hash
    }
    result:=make([][]string,0)//for give return
    for _,group:=range hashmap{
        result=append(result,group)
    }
    return result
}
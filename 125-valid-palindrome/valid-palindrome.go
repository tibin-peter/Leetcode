func isPalindrome(s string) bool {
    lowercased:=strings.ToLower(s)
    runed:=[]rune(lowercased)

    left,right:=0,len(runed)-1
    for left<right{
        if !unicode.IsLetter(runed[left]) && !unicode.IsDigit (runed[left]) {//avoid unwanted things on left
            left++
            continue
        }
        if !unicode.IsLetter(runed[right]) && !unicode.IsDigit (runed[right]) {//avoid unwanted things on right
            right--
            continue
        }
        if runed[left]!=runed[right]{
            return false
        }
          left++
          right--
    }
    return true
}
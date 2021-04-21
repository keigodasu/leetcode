func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    dicS := map[rune]int{}
    dicT := map[rune]int{}
   
    for _, v := range s {
        dicS[v]++
    }
    for _, v := range t {
        dicT[v]++
    }
    
    for k, v := range dicT {
        if vv, ok := dicS[k]; ok {
            if vv != v {
                return false
            }
        } else {
            return false
        }
    }
    
    return true
}

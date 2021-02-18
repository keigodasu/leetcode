func countConsistentStrings(allowed string, words []string) int {
    count := 0
    letters := map[rune]bool{}
    
    for _, v := range allowed {
        letters[v] = true
    }
    
    for _, v := range words {
        flag := false
        for _, vv := range v {
            if _, ok := letters[vv]; !ok {
                flag = true
                break
            }
        }
        
        if !flag {
            count++
        }
    }
    
    return count
}

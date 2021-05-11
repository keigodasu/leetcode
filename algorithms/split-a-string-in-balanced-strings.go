func balancedStringSplit(s string) int {
    stack := 0
    out := 0
    
    for _, v := range s {
        if v == 'R' {
            stack++
        } else {
            stack--
        }
        
        if stack == 0 {
            out++
        }
    }
    
    return out
}

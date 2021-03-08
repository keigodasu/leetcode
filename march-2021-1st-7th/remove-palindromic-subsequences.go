func removePalindromeSub(s string) int {
    if len(s) == 0 {
        return 0
    }

    reversed := ""
    for i := len(s) - 1; i >= 0; i-- {
        reversed += string(s[i])
    }
    
    if s == reversed {
        return 1
    } 
    
    return 2
}

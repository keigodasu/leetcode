func isStrobogrammatic(num string) bool {
    mapping := map[rune]rune{
        '0': '0',
        '1': '1',
        '6': '9',
        '8': '8',
        '9': '6',
    }
    
  
    reversed := ""
    for i := len(num) - 1; i >= 0; i-- {
        if _, ok := mapping[rune(num[i])]; ok {
            reversed += string(mapping[rune(num[i])]) 
        } else {
            return false
        }
    }
    
    return num == reversed
}

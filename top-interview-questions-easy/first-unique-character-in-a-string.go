func firstUniqChar(s string) int {
    dic := map[rune]int{}
    
    for _, v := range s {
        dic[v]++
    }
    
    for i, v := range s {
        if _, ok := dic[v]; ok {
            if dic[v] == 1 {
                return i
            }
        }
    }
 
    return -1
}

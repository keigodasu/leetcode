func hasAllCodes(s string, k int) bool {
    need := 1 << k
    dic := map[string]bool{}
    
    for i := 0; i <= len(s) - k; i++ {
        a := s[i:i+k]
        fmt.Println(a)
        if _, ok := dic[a]; !ok {
            dic[a] = true
            need--
            if need == 0 {
                return true
            }
        }
    }
    
    
    return false
}

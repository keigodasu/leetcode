//O(N^2)
func reverseString(s []byte)  {
    var temp []byte
    
    for i := len(s); i > 0; i-- {
        temp = append(temp, s[i-1])
    }
    
    for i := 0; i < len(s); i++ {
        s[i] = temp[i]
    }
}

func reverseString(s []byte)  {
    for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}


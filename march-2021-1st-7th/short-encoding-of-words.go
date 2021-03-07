func minimumLengthEncoding(words []string) int {
    list := []string{}
    dic := map[string]bool{}
    
    for i := 0; i < len(words); i++ {
        flag := false
        for j := 0; j < len(words); j++ {
            if i == j {
                continue
            }
            //完全一致
            if words[i] == words[j] {
                dic[words[i]] = true
            }
            //部分一致　
            if len(words[j]) >= len(words[i]) && words[i] == words[j][(len(words[j]) - len(words[i])):] {
                flag = true
                break
            }

        }
        if !flag {
            list = append(list, words[i]) 
        }
    }
   
    for k, _ := range dic {
        list = append(list, k) 
    }
   
    result := 0
    for _, v := range list {
        result += len(v) 
        result++
    }
    
    return result
}

func spellchecker(wordlist []string, queries []string) []string {
    var result []string
   
    words := map[string]string{}
    wordCap := map[string]string{}
    wordVow := map[string]string{}
    
    for _, v := range wordlist{
        words[v] = v
        
        if _, ok := wordCap[strings.ToLower(v)]; !ok {
            wordCap[strings.ToLower(v)] = v
        }
        
        if _, ok := wordVow[devowel(strings.ToLower(v))]; !ok {
            wordVow[devowel(strings.ToLower(v))] = v
        }
    }
    
    for _, v := range queries {
        
        if k, ok := words[v]; ok {
            result = append(result, k)    
            continue
        }
            
        if k, ok := wordCap[strings.ToLower(v)]; ok {
            result = append(result, k)    
            continue
        }

        if k, ok := wordVow[devowel(strings.ToLower(v))]; ok {
            result = append(result, k)    
            continue
        }
            
        result = append(result, "")    
    }
   
    return result
    
}

func devowel(word string) string {
    result := ""
    
    for _, v := range word {
        if isVowel(string(v)) {
            result += "*"
        } else {
            result += string(v)
        }
    }
    
    return result
}

func isVowel(letter string) bool {
    if letter == "a" || letter == "i" || letter == "u" || letter == "e" || letter == "o" {
        return true
    }
    
    return false
}

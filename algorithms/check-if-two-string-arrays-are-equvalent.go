func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    word01 := ""
    word02 := ""
   
    for _, v := range word1 {
        word01 += string(v)
    }
    for _, v := range word2 {
        word02 += string(v)
    }
    
    return word01 == word02
}

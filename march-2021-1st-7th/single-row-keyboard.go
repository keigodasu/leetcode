func calculateTime(keyboard string, word string) int {
    result := 0
    dic := map[rune]int{}
   
    for i, v := range keyboard {
        dic[v] = i
    }
   
    for i := 0; i < len(word) - 1; i++ {
        if i == 0 {
            result += dic[rune(word[i])]
        }
        result += int(math.Abs(float64(dic[rune(word[i])]) - float64(dic[rune(word[i+1])])))
    }
    
    return result
}

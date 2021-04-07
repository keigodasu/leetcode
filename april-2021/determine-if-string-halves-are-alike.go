func halvesAreAlike(s string) bool {
    lowS := strings.ToLower(s)
    halfPoint := (len(s) / 2)
    firstHalf := lowS[:halfPoint]
    secondHalf := lowS[halfPoint:]
    
    first := 0
    second := 0
    
    for _, v := range firstHalf {
        if v == 'a' || v == 'i' || v == 'u' || v == 'e' || v == 'o' {
            first++ 
        }
    }
    
    for _, v := range secondHalf {
        if v == 'a' || v == 'i' || v == 'u' || v == 'e' || v == 'o' {
            second++ 
        }
    }
    
    return first == second
    
}

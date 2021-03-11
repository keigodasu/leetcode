func plusOne(digits []int) []int {
    lastIndex := len(digits) - 1
    
    for i := lastIndex; i >= 0; i-- {
        if digits[i] == 9 {
            digits[i] = 0
        } else {
            digits[i]++
            return digits
        }
    }
   
    result := append([]int{1}, digits...)
    return result
}


func numberOfSteps (num int) int {
    count := 0
    current := num
    
    for current > 0 {
        count++
        if current % 2 == 0 {
            current = current / 2
        } else {
            current = current - 1
        }
    }
    
    return count
}

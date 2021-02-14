func maximumWealth(accounts [][]int) int {
    max := 0
    
    for _, v := range accounts {
        totalAmounts := 0
        for _, vv := range v {
            totalAmounts += vv
        }
        if totalAmounts >= max {
            max = totalAmounts
        }
    }
    
    return max
}

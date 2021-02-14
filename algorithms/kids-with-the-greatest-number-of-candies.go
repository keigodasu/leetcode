func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := findMax(candies)
    
    list := make([]bool, len(candies))
    for i, v := range candies {
        if v + extraCandies >= max {
            list[i] = true
        } else {
            list[i] = false
        }
    }
    
    return list
}

func findMax(candies []int) int {
    max := 0
    for _, v := range candies {
        if v >= max {
            max = v
        }
    }
    return max
}

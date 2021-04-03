func largestUniqueNumber(A []int) int {
    ordered := map[int]int{}
    
    for _, v := range A {
        ordered[v]++
    }
    
    max := 0
    for k, v := range ordered {
        if v == 1 && k > max {
            max = k
        } 
    }
    
   
    if max != 0 {
        return max
    } else {
        return -1
    }
}

func anagramMappings(A []int, B []int) []int {
    result := make([]int, len(A))
    
    for i, v := range A {
        for ii, vv := range B {
            if v == vv {
                result[i] = ii
            }
        }
    }
    
    return result
}

func anagramMappings(A []int, B []int) []int {
    result := make([]int, len(A))
   
    dicB := map[int]int{}
    for i, v := range B {
        dicB[v] = i
    }
    
    for i, v := range A {
        result[i] = dicB[v]
    }
    
    return result
}

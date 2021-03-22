func smallestCommonElement(mat [][]int) int {
    dic := map[int]int{}
    
    for _, v := range mat {
        for _, vv := range v {
            dic[vv]++
        }
    }
    
    for k, v := range dic {
        if v == len(mat) {
            return k
        }
    }
    
    return -1
}

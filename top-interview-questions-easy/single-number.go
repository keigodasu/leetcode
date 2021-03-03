func singleNumber(nums []int) int {
    result := 0
    sort.Ints(nums)
    
    dic := map[int]int{}
    
    for _, v := range nums {
        dic[v]++
    }
    

    for k, v := range dic {
        if v == 1 {
            result = k
        }
    }
    
    return result
}

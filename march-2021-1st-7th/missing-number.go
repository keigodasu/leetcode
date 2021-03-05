func missingNumber(nums []int) int {
    dic := map[int]bool{}
    
    for i := 0; i <= len(nums); i++ {
        dic[i] = false
    }
    
    for _, v := range nums {
       dic[v] = true
    }
    
    for k, v := range dic {
        if v == false {
            return k
        }
    }
    
    return 0
}
